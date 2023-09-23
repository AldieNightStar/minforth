package minforth

import (
	"errors"
	"fmt"
)

func notEnoughArgs(opname string, n int) error {
	return errors.New(fmt.Sprintf("Not enough args [%s]: %d", opname, n))
}
