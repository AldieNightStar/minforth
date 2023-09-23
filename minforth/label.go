package minforth

import (
	"errors"
	"fmt"
)

// Removes labels from code
func takeLabels(code *Code) (labs map[string]int) {
	labs = make(map[string]int)
	newops := []*operation{}

	var steps = 0
	for _, op := range code.Operations {
		if op.Type == OP_SPEC_DEF_LABEL {
			labs[op.Args[0]] = steps
			continue
		}
		steps += 1
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

func processLogicLabels(code *Code, labels map[string]int) ([]*operation, error) {
	newops := []*operation{}
	skips := 0
	for id, op := range code.Operations {
		// Skipping if needed
		if skips > 0 {
			skips -= 1
			continue
		}
		if isLogicSpecialType(op.Type) && op.Args[0] == "??" {
			next := getAt(code.Operations, id+1, nil)
			if next == nil {
				return nil, errors.New("After logic operator there are nothing")
			}
			if next.Type != OP_JUMP {
				return nil, errors.New("After logic operator there should be jump label")
			}
			jumpPos := next.Args[0]
			newops = append(newops, newOperation(op.Type, jumpPos, code.StackCell))
			newops = append(newops, newOperation(OP_NONE))

			// Skip next token
			skips = 1
		} else {
			newops = append(newops, op)
		}
	}
	return newops, nil
}

func isLogicSpecialType(t uint8) bool {
	return t == OP_SPEC_LT ||
		t == OP_SPEC_GT ||
		t == OP_SPEC_EQ ||
		t == OP_SPEC_GTE ||
		t == OP_SPEC_LTE ||
		t == OP_SPEC_NEQ
}
