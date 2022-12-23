package bondmachine

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/BondMachineHQ/BondMachine/pkg/simbox"
)

func (bmach *Bondmachine) WriteBMAPI(conf *Config, flavor string, iomaps *IOmap, extramods []ExtraModule, sbox *simbox.Simbox) error {

	var bmapiFlavor string
	var bmapiLanguage string
	var bmapiLibOutDir string
	var bmapiModOutDir string
	var bmapiAuxOutDir string
	var bmapiPackageName string
	var bmapiModuleName string

	var bmapiParams map[string]string

	// Extracting and check of BMAPI params
	for _, mod := range extramods {
		if mod.Get_Name() == "bmapi" {
			bmapiParams = mod.Get_Params().Params

			if val, ok := bmapiParams["bmapi_flavor"]; ok {
				bmapiFlavor = val
			} else {
				return errors.New("Missing bmapi flavor")
			}

			if val, ok := bmapiParams["bmapi_language"]; ok {
				bmapiLanguage = val
			} else {
				return errors.New("Missing bmapi language")
			}

			if val, ok := bmapiParams["bmapi_liboutdir"]; ok {
				bmapiLibOutDir = val
			} else {
				return errors.New("Missing bmapi liboutdir")
			}

			if val, ok := bmapiParams["bmapi_modoutdir"]; ok {
				bmapiModOutDir = val
			} else {
				return errors.New("Missing bmapi modoutdir")
			}

			if val, ok := bmapiParams["bmapi_auxoutdir"]; ok {
				bmapiAuxOutDir = val
			} else {
				return errors.New("Missing bmapi auxoutdir")
			}

			if val, ok := bmapiParams["bmapi_packagename"]; ok {
				bmapiPackageName = val
			} else {
				return errors.New("Missing bmapi packagename")
			}

			if val, ok := bmapiParams["bmapi_modulename"]; ok {
				bmapiModuleName = val
			} else {
				return errors.New("Missing bmapi modulename")
			}

			break
		}
	}

	switch bmapiFlavor {
	case "aximm":

		// This is the generation of the Linux kernel module
		if _, err := os.Stat(bmapiModOutDir); os.IsNotExist(err) {
			os.Mkdir(bmapiModOutDir, 0700)
		} else {
			return errors.New("BMAPI modoutdir already exists")
		}

		kmoddata := bmach.createBasicTemplateData()

		modFiles := make(map[string]string)
		modFiles["bm.c"] = moduleFilesBm

		for file, temp := range modFiles {
			t, err := template.New(file).Parse(temp)
			if err != nil {
				return err
			}

			f, err := os.Create(bmapiModOutDir + "/" + file)
			if err != nil {
				return err
			}

			err = t.Execute(f, kmoddata)
			if err != nil {
				return err
			}

			f.Close()
		}

		// This is the generation of the AXI auxiliary files
		if _, err := os.Stat(bmapiAuxOutDir); os.IsNotExist(err) {
			os.Mkdir(bmapiAuxOutDir, 0700)
		} else {
			return errors.New("BMAPI auxoutdir already exists")
		}

		auxdata := bmach.createBasicTemplateData()

		auxdata.Inputs = make([]string, 0)
		auxdata.Outputs = make([]string, 0)

		sortedKeys := make([]string, 0)
		for param, _ := range bmapiParams {
			sortedKeys = append(sortedKeys, param)
		}

		sort.Slice(sortedKeys, func(i, j int) bool {
			first := sortedKeys[i]
			second := sortedKeys[j]
			for {
				if len(first) == 0 || len(second) == 0 {
					return first < second
				} else {
					if first[0] != second[0] {
						return first < second
					} else {
						first = first[1:]
						second = second[1:]

						if numA, err := strconv.Atoi(first); err == nil {
							if numB, err := strconv.Atoi(second); err == nil {
								return numA < numB
							}
						}
					}
				}
			}
		})

		for _, param := range sortedKeys {
			if strings.HasPrefix(param, "assoc") {
				bmport := strings.Split(param, "_")[1]
				if strings.HasPrefix(bmport, "o") {
					auxdata.Outputs = append(auxdata.Outputs, "port_"+bmport)
				} else if strings.HasPrefix(bmport, "i") {
					auxdata.Inputs = append(auxdata.Inputs, "port_"+bmport)
				}
			}
		}

		auxFiles := make(map[string]string)
		auxFiles["axipatch.txt"] = auxfilesAXIPatch
		auxFiles["outregs.txt"] = auxfilesAXIOutRegs
		auxFiles["designexternal.txt"] = auxfilesDesignExternal
		auxFiles["designexternalinst.txt"] = auxfilesDesignExternalInst

		for file, temp := range auxFiles {
			t, err := template.New(file).Funcs(auxdata.funcmap).Parse(temp)
			if err != nil {
				return err
			}

			f, err := os.Create(bmapiAuxOutDir + "/" + file)
			if err != nil {
				return err
			}

			err = t.Execute(f, auxdata)
			if err != nil {
				return err
			}

			f.Close()
		}

		// Tivial aux files
		f, err := os.Create(bmapiAuxOutDir + "/axiregnum.txt")
		if err != nil {
			return err
		}
		f.Write([]byte(fmt.Sprintf("%d", len(auxdata.Inputs)+len(auxdata.Outputs)+4)))
		f.Close()

		switch bmapiLanguage {
		case "c":
			if _, err := os.Stat(bmapiLibOutDir); os.IsNotExist(err) {
				os.Mkdir(bmapiLibOutDir, 0700)
			} else {
				return errors.New("BMAPI liboutdir already exists")
			}

			// Compiling the data for the templates
			bmapidata := bmach.createBasicTemplateData()
			bmapidata.PackageName = bmapiPackageName
			var _ = bmapiModuleName // TODO TEMP
			cFiles := make(map[string]string)
			cFiles["Makefile"] = cFilesMakefile

			for file, temp := range cFiles {
				t, err := template.New(file).Parse(temp)
				if err != nil {
					return err
				}

				f, err := os.Create(bmapiLibOutDir + "/" + file)
				if err != nil {
					return err
				}

				err = t.Execute(f, bmapidata)
				if err != nil {
					return err
				}

				f.Close()
			}
		case "go":
			if _, err := os.Stat(bmapiLibOutDir); os.IsNotExist(err) {
				os.Mkdir(bmapiLibOutDir, 0700)
			} else {
				return errors.New("BMAPI liboutdir already exists")
			}

			// Compiling the data for the templates
			bmapidata := bmach.createBasicTemplateData()
			bmapidata.PackageName = bmapiPackageName
			var _ = bmapiModuleName // TODO TEMP
			apiFiles := make(map[string]string)
			apiFiles["bmapi.go"] = bmapi
			apiFiles["encoder.go"] = bmapiEncoder
			apiFiles["decoder.go"] = bmapiDecoder
			apiFiles["commands.go"] = bmapiCommands
			apiFiles["functions.go"] = bmapiFunctions
			apiFiles["go.mod"] = bmapigomod

			for file, temp := range apiFiles {
				t, err := template.New(file).Parse(temp)
				if err != nil {
					return err
				}

				f, err := os.Create(bmapiLibOutDir + "/" + file)
				if err != nil {
					return err
				}

				err = t.Execute(f, bmapidata)
				if err != nil {
					return err
				}

				f.Close()
			}
		}
	case "uartusb":
		switch bmapiLanguage {
		case "c":
			if _, err := os.Stat(bmapiLibOutDir); os.IsNotExist(err) {
				os.Mkdir(bmapiLibOutDir, 0700)
			} else {
				return errors.New("BMAPI liboutdir already exists")
			}

			// Compiling the data for the templates
			bmapidata := bmach.createBasicTemplateData()
			bmapidata.PackageName = bmapiPackageName
			var _ = bmapiModuleName // TODO TEMP
			cFiles := make(map[string]string)
			cFiles["Makefile"] = cFilesMakefile

			for file, temp := range cFiles {
				t, err := template.New(file).Parse(temp)
				if err != nil {
					return err
				}

				f, err := os.Create(bmapiLibOutDir + "/" + file)
				if err != nil {
					return err
				}

				err = t.Execute(f, bmapidata)
				if err != nil {
					return err
				}

				f.Close()
			}
		case "go":
			if _, err := os.Stat(bmapiLibOutDir); os.IsNotExist(err) {
				os.Mkdir(bmapiLibOutDir, 0700)
			} else {
				return errors.New("BMAPI liboutdir already exists")
			}

			// Compiling the data for the templates
			bmapidata := bmach.createBasicTemplateData()
			bmapidata.PackageName = bmapiPackageName
			var _ = bmapiModuleName // TODO TEMP
			apiFiles := make(map[string]string)
			apiFiles["bmapi.go"] = bmapi
			apiFiles["encoder.go"] = bmapiEncoder
			apiFiles["decoder.go"] = bmapiDecoder
			apiFiles["commands.go"] = bmapiCommands
			apiFiles["functions.go"] = bmapiFunctions
			apiFiles["go.mod"] = bmapigomod

			for file, temp := range apiFiles {
				t, err := template.New(file).Parse(temp)
				if err != nil {
					return err
				}

				f, err := os.Create(bmapiLibOutDir + "/" + file)
				if err != nil {
					return err
				}

				err = t.Execute(f, bmapidata)
				if err != nil {
					return err
				}

				f.Close()
			}
		default:
			return errors.New("unimplemented language")
		}
	default:
		return errors.New("unknown bmapi flavor")
	}

	return nil
}
