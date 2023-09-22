package minforth

import (
	"errors"
	"fmt"
)

// Removes labels from code
func takeLabels(code *Code) (labs map[string]int) {
	labs = make(map[string]int)
	newops := []*operation{}

	// Each special instruction has its own instruction set
	// For example stack operations is 2 instructions
	// Need to remember that to make precise labels work
	var instruction = 0

	// Everytime we delete something then all the elements shifts
	// Need to be up to that shifts to not miss after couple of them
	var shift = 0

	// Loop
	for _, op := range code.Operations {
		if op.Type == OP_SPEC_DEF_LABEL {
			labs[op.Args[0]] = instruction - shift
			shift += op.Type.Steps
			continue
		}
		instruction += op.Type.Steps
		newops = append(newops, op)
	}
	code.Operations = newops
	return labs
}

func processJumps(code *Code, labels map[string]int) ([]*operation, error) {
	newops := []*operation{}
	for _, op := range code.Operations {
		if op.Type == OP_SPEC_JUMP {
			labname := op.Args[0]
			labindex, labfound := labels[labname]
			if !labfound {
				return nil, errors.New("No label: " + labname)
			}
			newops = append(newops, newOperation(OP_JUMP, fmt.Sprint(labindex)))
		} else {
			newops = append(newops, op)
		}
	}
	return newops, nil
}
