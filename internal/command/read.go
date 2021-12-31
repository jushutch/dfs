package command

import (
	"fmt"
	"strings"
)

type ReadCommand struct {
	op   Operation
	path string
}

func (read ReadCommand) String() string {
	return fmt.Sprintf("%s %s", read.op, read.path)
}

func NewReadCommand(msg string) (Command, error) {
	parts := strings.Fields(msg)
	if len(parts) != 2 {
		return nil, fmt.Errorf("incorrect number of arguments to read")
	}
	return &ReadCommand{READ, parts[1]}, nil
}
