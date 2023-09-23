package minforth

func (o *operation) logicJump(condition string) (newops []*operation, reexpand bool, err error) {
	if o.HasAllArgs(2) {
		// Here are we are combining stack pops and jump logic to special position if conditions are met
		return []*operation{
			newOperation(OP_SPEC_POP, "VAL2", o.Args[1]),
			newOperation(OP_SPEC_POP, "VAL1", o.Args[1]),
			newOperation(OP_JUMP_COND, o.Args[0], condition, "VAL1", "VAL2"),
		}, true, nil
	} else {
		return nil, false, notEnoughArgs("jump if "+condition, 2)
	}
}

func (o *operation) expand() (mewops []*operation, reexpand bool, err error) {
	if o.Type == OP_SPEC_PUSH {
		// Args: value cell_name
		if o.HasAllArgs(2) {
			return []*operation{
				newOperation(OP_CELL_WRITE, o.Args[0], o.Args[1], "SP"),
				newOperation(OP_ADD, "SP", "SP", "1"),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("push", 2)
		}
	} else if o.Type == OP_SPEC_POP {
		// o.Args: result cell_name
		if o.HasAllArgs(2) {
			return []*operation{
				newOperation(OP_SUB, "SP", "SP", "1"),
				newOperation(OP_CELL_READ, o.Args[0], o.Args[1], "SP"),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("pop", 2)
		}
	} else if o.Type == OP_SPEC_DUPE {
		// Args: cell_name
		if o.HasAllArgs(1) {
			return []*operation{
				newOperation(OP_SPEC_POP, "VAL1", o.Args[0]),
				newOperation(OP_SPEC_PUSH, "VAL1", o.Args[0]),
				newOperation(OP_SPEC_PUSH, "VAL1", o.Args[0]),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("dupe", 1)
		}
	} else if o.Type == OP_SPEC_DROP {
		// Args: cell_name
		if o.HasAllArgs(1) {
			return []*operation{
				newOperation(OP_SPEC_POP, "VAL1", o.Args[0]),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("drop", 1)
		}
	} else if o.Type == OP_SPEC_SET_VAR {
		// Args: var_name cell_name
		if o.HasAllArgs(2) {
			return []*operation{
				newOperation(OP_SPEC_POP, o.Args[0], o.Args[1]),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("set var", 1)
		}
	} else if o.Type == OP_SPEC_GET_VAR {
		// Args: var_name cell_name
		if o.HasAllArgs(2) {
			return []*operation{
				newOperation(OP_SPEC_PUSH, o.Args[0], o.Args[1]),
			}, true, nil
		} else {
			return nil, false, notEnoughArgs("get var", 1)
		}
	} else if o.Type == OP_SPEC_LT {
		return o.logicJump("lessThan")
	} else if o.Type == OP_SPEC_EQ {
		return o.logicJump("equal")
	} else if o.Type == OP_SPEC_GT {
		return o.logicJump("greaterThan")
	} else if o.Type == OP_SPEC_LTE {
		return o.logicJump("lessThanEq")
	} else if o.Type == OP_SPEC_GTE {
		return o.logicJump("greaterThanEq")
	} else if o.Type == OP_SPEC_NEQ {
		return o.logicJump("notEqual")
	}
	return []*operation{o}, false, nil
}

func expandAll(ops []*operation) (newops []*operation, reexpand bool, err error) {
	// Reusable
	var expandedOps []*operation
	var _reexpand = false

	for _, op := range ops {
		expandedOps, _reexpand, err = op.expand()
		// Once reexpand is true then it's true until end
		if _reexpand {
			reexpand = true
		}
		if err != nil {
			return nil, false, err
		}
		for _, expandedOp := range expandedOps {
			newops = append(newops, expandedOp)
		}
	}

	return newops, reexpand, nil
}

func expandAllWithReexpand(ops []*operation) (newops []*operation, err error) {
	var reexpand = true
	newops = ops
	for reexpand {
		newops, reexpand, err = expandAll(newops)
		if err != nil {
			return nil, err
		}
	}
	return newops, nil
}
