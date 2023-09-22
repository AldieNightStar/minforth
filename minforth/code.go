package minforth

import "strings"

type Code struct {
	StackCell   string
	MessageCell string
	Operations  []*operation
}

func newCode(stackCell, messageCell string) *Code {
	code := &Code{
		StackCell:   stackCell,
		MessageCell: messageCell,
		Operations:  make([]*operation, 0, 8),
	}
	code.Add(newOperation(OP_SET, "SP", "0"))
	return code
}

func (c *Code) Add(op *operation) {
	c.Operations = append(c.Operations, op)
}

func (c *Code) String() (string, error) {
	var arr []string
	for _, op := range c.Operations {
		result, err := op.String()
		if err != nil {
			return "", err
		}
		arr = append(arr, result)
	}
	return strings.Join(arr, "\n"), nil
}
