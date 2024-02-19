package bmqsim

import (
	"fmt"
	"strings"

	"github.com/BondMachineHQ/BondMachine/pkg/bmline"
	"github.com/BondMachineHQ/BondMachine/pkg/bmmatrix"
)

type BmQSimulator struct {
	verbose  bool
	debug    bool
	qbits    []string
	qbitsNum map[string]int
}

// BmQSimulatorInit initializes the BmQSimulator
func (sim *BmQSimulator) BmQSimulatorInit() {
	sim.verbose = false
	sim.debug = false
	sim.qbits = make([]string, 0)
	sim.qbitsNum = make(map[string]int)
}

func (sim *BmQSimulator) Dump() string {
	return fmt.Sprintf("BmQSimulator: verbose=%t, debug=%t, qbits=%v, qbitsNum=%v", sim.verbose, sim.debug, sim.qbits, sim.qbitsNum)
}

func (sim *BmQSimulator) SetVerbose() {
	sim.verbose = true
}

func (sim *BmQSimulator) SetDebug() {
	sim.debug = true
}

// QasmToBmMatrices converts a QASM file to a list of BmMatrixSquareComplex, the input is a BasmBody with all the metadata and the list of quantum instructions
func (sim *BmQSimulator) QasmToBmMatrices(qasm *bmline.BasmBody) ([]*bmmatrix.BmMatrixSquareComplex, error) {

	result := make([]*bmmatrix.BmMatrixSquareComplex, 0)

	// Get the qbits and their names
	qbits := qasm.GetMeta("qbits")

	if qbits == "" {
		return nil, fmt.Errorf("no qbits defined")
	}

	for i, qbit := range strings.Split(qbits, ":") {
		if _, ok := sim.qbitsNum[qbit]; ok {
			return nil, fmt.Errorf("qbit %s already defined", qbit)
		} else {
			sim.qbits = append(sim.qbits, qbit)
			sim.qbitsNum[qbit] = i
		}
	}

	curOp := make([]*bmline.BasmLine, 0)
	curQBits := make(map[int]struct{})

	for i, line := range qasm.Lines {
		op := line.Operation.GetValue()

		if sim.debug {
			fmt.Printf("Processing line %d: %s\n", i, line.String())
		}

		// Check if the operation is ready to form a matrix
		nextOp := false

		// Include the qbits in the operation to the currQbits map
		for _, arg := range line.Elements {
			argName := arg.GetValue()
			// Check if the argument is a qbit, otherwise ignore it
			if qbitN, ok := sim.qbitsNum[argName]; ok {
				if _, ok := curQBits[qbitN]; ok {
					nextOp = true
					break
				} else {
					curQBits[sim.qbitsNum[argName]] = struct{}{}
				}
			}

		}

		if op == "nextop" {
			nextOp = true
		}

		singleLast := false

		if i == len(qasm.Lines)-1 {
			// If the last line is not already a nextOp lets put it in current operation, otherwise we will set singleLast to true
			// and process it alone later on
			if !nextOp {
				curOp = append(curOp, line)
			} else {
				singleLast = true
			}
			nextOp = true
		}

		// If the operation is ready to form a matrix, create the matrix
		if nextOp {
			// Create the matrix
			if len(curOp) > 0 {
				// Create the matrix
				if m, err := sim.BmMatrixFromOperation(curOp); err != nil {
					return nil, fmt.Errorf("error creating matrix from operation: %v", err)
				} else {
					if m != nil {
						result = append(result, m)
					}
				}
			}

			// Reset the operation and the qbits
			curOp = make([]*bmline.BasmLine, 0)
			curQBits = make(map[int]struct{})
		}

		// If the last line is not already been added to the last matrix, lets put it alone in a new matrix
		// Otherwise (not the last or already done) we will put it in the current operations list and set the involved qbits into the currQbits map
		if singleLast {
			curOp = append(curOp, line)
			// Create the matrix
			if m, err := sim.BmMatrixFromOperation(curOp); err != nil {
				return nil, fmt.Errorf("error creating matrix from operation: %v", err)
			} else {
				if m != nil {
					result = append(result, m)
				}
			}

		} else {
			curOp = append(curOp, line)

			// Include the qbits in the operation to the currQbits map
			for _, arg := range line.Elements {
				argName := arg.GetValue()
				// Check if the argument is a qbit, otherwise ignore it
				if qbitN, ok := sim.qbitsNum[argName]; ok {
					curQBits[qbitN] = struct{}{}
				}

			}

		}
	}
	return result, nil
}

type swap struct {
	s1 int
	s2 int
}

func (sim *BmQSimulator) BmMatrixFromOperation(op []*bmline.BasmLine) (*bmmatrix.BmMatrixSquareComplex, error) {
	// Let prepare the matrix that will be tensor-producted to form the final matrix
	var result *bmmatrix.BmMatrixSquareComplex

	swaps := make([]swap, 0)
	// loop over the qbits each qbit will have only one (or zero) operation within the op list
	// Every line where the qbits in not in sequence will be reordered swapping the qbits
	// and swapped back after the matrix is created

	localQBits := make([]string, len(sim.qbits))
	copy(localQBits, sim.qbits)

	for q := 0; q < len(localQBits); q++ {
		qbit := localQBits[q]
		// Find the operation for the qbit
		found := false
		fundLine := -1
		for i, line := range op {
			for _, arg := range line.Elements {
				argName := arg.GetValue()
				if argName == qbit {
					found = true
					fundLine = i
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			// No operation for the qbit, lets add an identity matrix
			// Create the identity matrix
			ident := bmmatrix.IdentityComplex(2)
			if result == nil {
				result = ident
			} else {
				result = bmmatrix.TensorProductComplex(result, ident)
			}
		} else {
			argNumQBits := len(op[fundLine].Elements)
			if argNumQBits == 1 {
				// Single qbit operation
				// Create the matrix
				if matrix, err := sim.MatrixFromOp(op[fundLine].Operation.GetValue()); err != nil {
					return nil, fmt.Errorf("error creating matrix from operation: %v", err)
				} else {
					if result == nil {
						result = matrix
					} else {
						result = bmmatrix.TensorProductComplex(result, matrix)
					}
				}
			} else {
				// Multi qbit operation
				// The order of the qbits in the operation is important, we need to reorder the qbits in the operation
				// if they are not in sequence by swapping the qbits and then swapping them back after the matrix is created

				localOrder := make([]int, argNumQBits)
				for i, arg := range op[fundLine].Elements {
					argName := arg.GetValue()
					localOrder[i] = sim.qbitsNum[argName]
				}

				fmt.Println(localOrder, q)

				for i, lq := range localOrder {
					if lq != q {
						// Swap the qbits
						localQBits[q], localQBits[lq] = localQBits[lq], localQBits[q]
						// Add the swap to the list
						swaps = append(swaps, swap{q, lq})
						if sim.debug {
							fmt.Printf("Swapping qbits %d and %d\n", q, lq)
						}
						// Swap the localOrder if needed
						for j, lq2 := range localOrder {
							if lq2 == q {
								localOrder[j] = lq
							} else if lq2 == lq {
								localOrder[j] = q
							}
						}
						if sim.debug {
							fmt.Println("newLocalOrder:", localOrder)
						}

					}
					if i != len(localOrder)-1 {
						q++
					}
				}

				// Create the matrix
				if matrix, err := sim.MatrixFromOp(op[fundLine].Operation.GetValue()); err != nil {
					return nil, fmt.Errorf("error creating matrix from operation: %v", err)
				} else {
					if result == nil {
						result = matrix
					} else {
						result = bmmatrix.TensorProductComplex(result, matrix)
					}
				}
			}
		}
	}

	if sim.debug {
		fmt.Println("swaps:", swaps)
		swaps2baseSwaps(swap{0, 0}, len(sim.qbits))
	}

	return result, nil
}

func swaps2baseSwaps(s swap, n int) []swap {
	baseNum := uint64(1 << n)
	fmt.Println("baseNum:", baseNum)
	// TODO: implement
	return nil
}

func (sim *BmQSimulator) MatrixFromOp(op string) (*bmmatrix.BmMatrixSquareComplex, error) {
	switch op {
	case "h", "H":
		return bmmatrix.Hadamard(), nil
	case "x", "X":
		return bmmatrix.PauliX(), nil
	case "y", "Y":
		return bmmatrix.PauliY(), nil
	case "z", "Z":
		return bmmatrix.PauliZ(), nil
	case "cx", "CX":
		return bmmatrix.CNot(), nil
	case "zero", "ZERO":
	// Ignore the zero operation
	default:
		return nil, fmt.Errorf("unknown operation %s", op)
	}
	return nil, nil
}
