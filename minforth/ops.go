package minforth

import (
	"errors"
	"fmt"
	"strings"
)

func notEnoughArgs(opname string, n int) error {
	return errors.New(fmt.Sprintf("Not enough args [%s]: %d", opname, n))
}

type operationType struct {
	// How many instructions used per this opertaion
	Steps int
}

var (
	// Args: message
	OP_PRINT = &operationType{1}

	// Args: cell_name
	OP_PRINT_FLUSH = &operationType{1}

	// Args: name value
	OP_SET = &operationType{1}

	// Args: result a b
	OP_ADD = &operationType{1}

	// Args: result a b
	OP_SUB = &operationType{1}

	// Args: result a b
	OP_MUL = &operationType{1}

	// Args: result a b
	OP_DIV = &operationType{1}

	// Args: result cell_name at
	OP_CELL_READ = &operationType{1}

	// Args: value cell_name at
	OP_CELL_WRITE = &operationType{1}

	// Args: seconds
	OP_WAIT = &operationType{1}

	// Args: pos
	OP_JUMP = &operationType{1}

	// =====================
	// Special Operations
	// =====================

	// Args: value cell_name
	OP_SPEC_PUSH = &operationType{2}

	// Args: result cell_name
	OP_SPEC_POP = &operationType{2}

	// Args: cell_name
	OP_SPEC_DUPE = &operationType{6}

	// Args: cell_name
	OP_SPEC_DROP = &operationType{2}

	// Args: label_name
	OP_SPEC_DEF_LABEL = &operationType{1}

	// Args: label_name
	OP_SPEC_JUMP = &operationType{1}

	// Args: var_name cell_name
	OP_SPEC_SET_VAR = &operationType{2}

	// Args: var_name cell_name
	OP_SPEC_GET_VAR = &operationType{2}

	// Args: jump_pos cell_name
	OP_SPEC_LT = &operationType{5}

	// Args: jump_pos cell_name
	OP_SPEC_GT = &operationType{5}

	// Args: jump_pos cell_name
	OP_SPEC_EQ = &operationType{5}

	// Args: jump_pos cell_name
	OP_SPEC_NEQ = &operationType{5}

	// Args: jump_pos cell_name
	OP_SPEC_GTE = &operationType{5}

	// Args: jump_pos cell_name
	OP_SPEC_LTE = &operationType{5}

	// Args: 0
	OP_NONE = &operationType{0}
)

type operation struct {
	Type *operationType
	Args []string
}

func newOperation(t *operationType, args ...string) *operation {
	return &operation{
		Type: t,
		Args: args,
	}
}

func (o *operation) HasAllArgs(n int) bool {
	return len(o.Args) >= n
}

func (o *operation) joinedArgs() string { return strings.Join(o.Args, " ") }

func (o *operation) simpleOperation(name string, argcount int) (string, error) {
	if o.HasAllArgs(argcount) {
		return fmt.Sprintf("%s %s", name, o.joinedArgs()), nil
	} else {
		return "", notEnoughArgs(name, argcount)
	}
}

func (o *operation) logicJump(condition string) (string, error) {
	if o.HasAllArgs(2) {
		// Here are we are combining stack pops and jump logic to special position if conditions are met
		result, err := o.combine(
			newOperation(OP_SPEC_POP, "VAL2", o.Args[1]),
			newOperation(OP_SPEC_POP, "VAL1", o.Args[1]),
		)
		if err != nil {
			return "", err
		}
		return result + "\n" + fmt.Sprintf("jump %s %s VAL1 VAL2", o.Args[0], condition), nil
	} else {
		return "", notEnoughArgs("jump if "+condition, 2)
	}
}

func (o *operation) combine(ops ...*operation) (string, error) {
	array := []string{}
	for _, op := range ops {
		result, err := op.String()
		if err != nil {
			return "", err
		}
		array = append(array, result)
	}
	return strings.Join(array, "\n"), nil
}

func (o *operation) String() (string, error) {
	// Checks
	if o.Type == OP_PRINT {
		return o.simpleOperation("print", 1)
	} else if o.Type == OP_PRINT_FLUSH {
		return o.simpleOperation("printflush", 1)
	} else if o.Type == OP_SET {
		return o.simpleOperation("set", 2)
	} else if o.Type == OP_ADD {
		return o.simpleOperation("op add", 3)
	} else if o.Type == OP_SUB {
		return o.simpleOperation("op sub", 3)
	} else if o.Type == OP_MUL {
		return o.simpleOperation("op mul", 3)
	} else if o.Type == OP_DIV {
		return o.simpleOperation("op div", 3)
	} else if o.Type == OP_CELL_READ {
		return o.simpleOperation("read", 3)
	} else if o.Type == OP_CELL_WRITE {
		return o.simpleOperation("write", 3)
	} else if o.Type == OP_WAIT {
		return o.simpleOperation("wait", 1)
	} else if o.Type == OP_JUMP {
		// o.Args: pos
		if o.HasAllArgs(1) {
			return fmt.Sprintf("jump %s always ___ 0", o.Args[0]), nil
		} else {
			return "", notEnoughArgs("jump", 1)
		}
	} else if o.Type == OP_SPEC_PUSH {
		// o.Args: value cell_name
		if o.HasAllArgs(2) {
			return o.combine(
				newOperation(OP_CELL_WRITE, o.Args[0], o.Args[1], "SP"),
				newOperation(OP_ADD, "SP", "SP", "1"),
			)
		} else {
			return "", notEnoughArgs("push", 2)
		}
	} else if o.Type == OP_SPEC_POP {
		// o.Args: result cell_name
		if o.HasAllArgs(2) {
			return o.combine(
				newOperation(OP_SUB, "SP", "SP", "1"),
				newOperation(OP_CELL_READ, o.Args[0], o.Args[1], "SP"),
			)
		} else {
			return "", notEnoughArgs("pop", 2)
		}
	} else if o.Type == OP_SPEC_DUPE {
		if o.HasAllArgs(1) {
			return o.combine(
				newOperation(OP_SPEC_POP, "VAL1", o.Args[0]),
				newOperation(OP_SPEC_PUSH, "VAL1", o.Args[0]),
				newOperation(OP_SPEC_PUSH, "VAL1", o.Args[0]),
			)
		} else {
			return "", notEnoughArgs("dupe", 1)
		}
	} else if o.Type == OP_SPEC_DROP {
		if o.HasAllArgs(1) {
			return o.combine(
				newOperation(OP_SPEC_POP, "VAL1", o.Args[0]),
			)
		} else {
			return "", notEnoughArgs("drop", 1)
		}
	} else if o.Type == OP_SPEC_DEF_LABEL {
		// Return empty render. It does nothing
		return "noop", nil
	} else if o.Type == OP_SPEC_SET_VAR {
		if o.HasAllArgs(2) {
			return o.combine(
				newOperation(OP_SPEC_POP, o.Args[0], o.Args[1]),
			)
		} else {
			return "", notEnoughArgs("set var", 1)
		}
	} else if o.Type == OP_SPEC_GET_VAR {
		if o.HasAllArgs(2) {
			return o.combine(
				newOperation(OP_SPEC_PUSH, o.Args[0], o.Args[1]),
			)
		} else {
			return "", notEnoughArgs("get var", 1)
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
	} else if o.Type == OP_NONE {
		return "noop", nil
	}
	return "", errors.New("Unknown operation")
}
