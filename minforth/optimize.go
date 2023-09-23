package minforth

func optimize(code *Code, times int) {
	for i := 0; i < times; i++ {
		code.Operations = optimizeSetRead(code)
	}
}

// This code is working out of 'optimize' method
// Before expand was made
func optimizePushPop(code *Code) []*operation {
	newops := []*operation{}
	skips := 0
	for id, op := range code.Operations {
		// Skip couple of elements if needed
		if skips > 0 {
			skips -= 1
			continue
		}

		// Remove push and pops
		// If we have PUSH then POP - we replace to simple SET VAL1
		if op.Type == OP_SPEC_PUSH {
			next := getAt(code.Operations, id+1, nil)
			if next != nil {
				if next.Type == OP_SPEC_POP {
					newops = append(newops, newOperation(OP_SET, "VAL1", op.Args[0]))
					skips = 1 // Skip one more
					continue
				}
			}
		}
		newops = append(newops, op)
	}
	return newops
}

func optimizeSetRead(code *Code) []*operation {
	newops := []*operation{}
	skips := 0
	for id, op := range code.Operations {
		// Skip couple of elements if needed
		if skips > 0 {
			skips -= 1
			continue
		}

		// If current operation is setting VAL1 and second using it
		// Then make next one use that value directly
		if op.Type == OP_SET && op.Args[0] == "VAL1" {
			next := getAt(code.Operations, id+1, nil)
			if next != nil {
				if isUsingValue(next, "VAL1") {
					next = opReplaceVal(next, "VAL1", op.Args[1])
					newops = append(newops, next)
					skips = 1 // Skip one more
					continue
				}
			}
		}
		newops = append(newops, op)
	}
	return newops
}

func isUsingValue(op *operation, value string) bool {
	for _, arg := range op.Args {
		if arg == value {
			return true
		}
	}
	return false
}

func opReplaceVal(op *operation, oldValue string, value string) *operation {
	for id, arg := range op.Args {
		if arg == oldValue {
			op.Args[id] = value
		}
	}
	return op
}
