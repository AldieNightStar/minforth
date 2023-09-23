package minforth

import (
	"errors"
	"fmt"
	"strings"
)

type operation struct {
	Type uint8
	Args []string
}

func newOperation(t uint8, args ...string) *operation {
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
		if o.HasAllArgs(1) {
			return fmt.Sprintf("jump %s always", o.Args[0]), nil
		} else {
			return "", notEnoughArgs("jump", 1)
		}
	} else if o.Type == OP_JUMP_COND {
		return o.simpleOperation("jump", 4)
	} else if o.Type == OP_NONE {
		return "noop", nil
	} else if o.Type == OP_CONTROL {
		return o.simpleOperation("control", 3)
	} else if o.Type == OP_SENSOR {
		return o.simpleOperation("sensor", 3)
	}
	return "", errors.New("Unknown operation")
}
