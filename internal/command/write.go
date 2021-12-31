package command

import (
	"fmt"
)

type WriteCommand struct {
	op        Operation
	fileBlock int
	path      string
	data      []byte
}

func (write WriteCommand) String() string {
	return fmt.Sprintf("%s %q %d %s", write.op, write.path, write.fileBlock, write.data)
}

// func NewWriteCommand(msg string) (Command, error) {
// 	parts := strings.Fields(msg)
// 	if len(parts) != 2 {
// 		return nil, fmt.Errorf("incorrect number of arguments to write")
// 	}
// 	os.Args = parts
// 	var block *int = flag.Int("block", -1, "The data block to write to. The default appends to the end of the file")
// 	flag.Parse()
// 	fmt.Print(block)
// 	flag
// }
