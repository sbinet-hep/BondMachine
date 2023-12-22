package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/BondMachineHQ/BondMachine/pkg/bmgraph"
	"github.com/BondMachineHQ/BondMachine/pkg/bminfo"
	"github.com/BondMachineHQ/BondMachine/pkg/bmnumbers"
	graphviz "github.com/goccy/go-graphviz"
)

var verbose = flag.Bool("v", false, "Verbose")
var debug = flag.Bool("d", false, "Debug")

var registerSize = flag.Int("register-size", 32, "Number of bits per register (n-bit)")
var dataType = flag.String("data-type", "float32", "bmnumbers data types")

var saveBasm = flag.String("save-basm", "", "Create a basm file")

var neuronLibPath = flag.String("neuron-lib-path", "", "Path to the neuron library to use")

var graphFile = flag.String("graph-file", "", "Graph (DOT)")
var configFile = flag.String("config-file", "", "JSON description of the net configuration")
var bmInfoFile = flag.String("bminfo-file", "", "JSON description of the BondMachine abstraction")

var iomode = flag.String("io-mode", "async", "IO mode: async, sync")

func init() {
	flag.Parse()
	if *saveBasm == "" {
		*saveBasm = "out.basm"
	}
}

func main() {
	g := graphviz.New()

	var data []byte

	// Load a graph from a file
	if *graphFile != "" {
		var err error
		data, err = os.ReadFile(*graphFile)
		if err != nil {
			panic(err)
		}
	} else {
		log.Fatalln("No graph file specified")
	}

	graph, err := graphviz.ParseBytes(data)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := graph.Close(); err != nil {
			panic(err)
		}
		g.Close()
	}()

	// Create the config struct
	config := new(bmgraph.Config)

	// Load net from a JSON file the configuration
	if *configFile != "" {
		if netFileJSON, err := os.ReadFile(*configFile); err == nil {
			if err := json.Unmarshal(netFileJSON, config); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		config.Debug = *debug
		config.Verbose = *verbose
		config.Params = make(map[string]string)
	}

	// Load or create the Info file
	config.BMinfo = new(bminfo.BMinfo)

	if *bmInfoFile != "" {
		if bmInfoJSON, err := os.ReadFile(*bmInfoFile); err == nil {
			if err := json.Unmarshal(bmInfoJSON, config.BMinfo); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	if config.Params == nil {
		config.Params = make(map[string]string)
	}
	if config.List == nil {
		config.List = make(map[string]string)
	}
	if config.Pruned == nil {
		config.Pruned = make([]string, 0)
	}

	if *neuronLibPath != "" {
		config.NeuronLibPath = *neuronLibPath
	} else {
		// panic("No neuron library path specified")
	}

	if *dataType != "" {
		found := false
		for _, tpy := range bmnumbers.AllTypes {
			if tpy.GetName() == *dataType {
				for opType, opName := range tpy.ShowInstructions() {
					config.Params[opType] = opName
				}
				config.DataType = *dataType
				config.TypePrefix = tpy.ShowPrefix()
				config.Params["typeprefix"] = tpy.ShowPrefix()
				found = true
				break
			}
		}
		if !found {
			if created, err := bmnumbers.EventuallyCreateType(*dataType, nil); err == nil {
				if created {
					for _, tpy := range bmnumbers.AllTypes {
						if tpy.GetName() == *dataType {
							for opType, opName := range tpy.ShowInstructions() {
								config.Params[opType] = opName
							}
							config.DataType = *dataType
							config.TypePrefix = tpy.ShowPrefix()
							config.Params["typeprefix"] = tpy.ShowPrefix()
							break
						}
					}
				} else {
					panic("Unknown data type")
				}

			} else {
				panic(err)
			}
		}
	} else {
		if config.DataType == "" {
			panic("No data type specified")
		}
	}

	bg := new(bmgraph.Graph)
	bg.Graph = graph

	if *saveBasm != "" {
		if basmFile, err := bg.WriteBasm(); err == nil {
			os.WriteFile(*saveBasm, []byte(basmFile), 0644)
		} else {
			panic(err)
		}
	}

	if *bmInfoFile != "" {
		// Write the info file
		if bmInfoFileJSON, err := json.MarshalIndent(config.BMinfo, "", "  "); err == nil {
			os.WriteFile(*bmInfoFile, bmInfoFileJSON, 0644)
		} else {
			panic(err)
		}
	}

	// Remove the info file from the config prior to saving it
	config.BMinfo = nil
	if *configFile != "" {
		// Write the eventually updated config file
		if configFileJSON, err := json.MarshalIndent(config, "", "  "); err == nil {
			os.WriteFile(*configFile, configFileJSON, 0644)
		} else {
			panic(err)
		}
	}

}
