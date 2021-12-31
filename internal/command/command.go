package command

import (
	"fmt"
	"strings"
)

// Enum for supported commands
type Operation string

const (
	READ   Operation = "read"
	WRITE  Operation = "write"
	CREATE Operation = "create"
	DELETE Operation = "delete"
	TREE   Operation = "tree"
	USER   Operation = "user"
)

type Command interface {
	String() string
}

// Create a new and valid Command from any message
func NewCommand(msg string) (Command, error) {
	parts := strings.Fields(msg)
	op := strings.ToLower(parts[0])
	switch op {
	case "read":
		return NewReadCommand(msg)
	// case "write":
	// 	return &Command{WRITE, ' '}, nil
	// case "create":
	// 	if len(parts) != 3 {
	// 		return nil, fmt.Errorf("invalid create command")
	// 	}
	// 	if parts[1] != "f" && parts[1] != "d" {
	// 		return nil, fmt.Errorf("invalid create type")
	// 	}
	// 	return &Command{CREATE, parts[1][0]}, nil
	// case "delete":
	// 	return &Command{DELETE, ' '}, nil
	default:
		return nil, fmt.Errorf("\"%s\" is an unrecognized file system operation", op)
	}
}
