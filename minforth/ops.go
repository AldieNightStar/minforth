package minforth

import (
	"errors"
	"fmt"
	"strings"
)

func notEnoughArgs(opname string, n int) error {
	return errors.New(fmt.Sprintf("Not enough args [%s]: %d", opname, n))
}

const (
	// Args: message
	OP_PRINT = iota

	// Args: cell_name
	OP_PRINT_FLUSH

	// Args: name value
	OP_SET

	// Args: result a b
	OP_ADD

	// Args: result a b
	OP_SUB

	// Args: result a b
	OP_MUL

	// Args: result a b
	OP_DIV

	// Args: result cell_name at
	OP_CELL_READ

	// Args: value cell_name at
	OP_CELL_WRITE

	// Args: seconds
	OP_WAIT

	// Args: pos
	OP_JUMP

	// =====================
	// Special Operations
	// =====================

	// Args: value cell_name
	OP_SPEC_PUSH

	// Args: result cell_name
	OP_SPEC_POP

	// Args: cell_name
	OP_SPEC_DUPE

	// Args: cell_name
	OP_SPEC_DROP

	// Args: label_name
	OP_SPEC_DEF_LABEL

	// Args: label_name
	OP_SPEC_JUMP

	// Args: var_name cell_name
	OP_SPEC_SET_VAR

	// Args: var_name cell_name
	OP_SPEC_GET_VAR
)

type operation struct {
	Type int
	Args []string
}

func newOperation(typ int, args ...string) *operation {
	return &operation{
		Type: typ,
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
	}
	return "", errors.New("Unknown operation")
}
