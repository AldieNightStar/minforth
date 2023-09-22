package minforth

import "fmt"

type Code struct {
	StackCell string
	Lines     []string
}

func (c *Code) GetLine() int {
	return len(c.Lines)
}

func (c *Code) SetMem(name string, value int) {
	c.Lines = append(
		c.Lines,
		fmt.Sprintf("set %s %d", name, value),
	)
}

// op could be: add, sub, div, mul
func (c *Code) Operate(name string, op string, val2 string) {
	c.Lines = append(
		c.Lines,
		fmt.Sprintf("op %s %s %s", op, name, val2),
	)
}

func (c *Code) CellWrite(name string, id int) {
	c.Lines = append(
		c.Lines,
		fmt.Sprintf("write %s %s %d", name, c.StackCell, id),
	)
}

func (c *Code) CellRead(name string, id int) {
	c.Lines = append(
		c.Lines,
		fmt.Sprintf("read %s %s %d", name, c.StackCell, id),
	)
}

func (c *Code) PushStack(value int) {
	c.Lines = append(
		c.Lines,
		fmt.Sprint("op"),
	)
}

func (c *Code) Jump(line int) {
	c.Lines = append(
		c.Lines,
		fmt.Sprintf("jump %d always false", line),
	)
}

func NewCode(stackCell string) *Code {
	return &Code{
		StackCell: stackCell,
		Lines:     make([]string, 0, 32),
	}
}
