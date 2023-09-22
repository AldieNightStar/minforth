package minforth

import (
	"errors"
	"fmt"
)

// Removes labels from code
func takeLabels(code *Code) (labs map[string]int) {
	labs = make(map[string]int)
	newops := []*operation{}
	var shift = 0
	for id, op := range code.Operations {
		if op.Type == OP_SPEC_DEF_LABEL {
			labs[op.Args[0]] = id - shift
			shift += 1
			continue
		}
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
