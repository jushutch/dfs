package command

import "testing"

func TestInvalidMessage(t *testing.T) {
	var tests = []string{
		"readd",
		"writ",
		"crete",
		"create file.txt",
		"create a file.txt",
		"del",
	}

	for _, message := range tests {
		if _, err := NewCommand(message); err == nil {
			t.Errorf("Test Failed: \"%s\" inputted, error expected", message)
		}
	}
}

func TestValidMessage(t *testing.T) {
	var tests = []string{
		"read file.txt",
		"write file.txt \"string to be written\"",
		"write file.txt -b 3 \"string to be written\"",
		"write -b 3 file.txt \"string to be written\"",
		"create f file.txt",
		"create d file.txt",
		"create f /dir/file.txt",
		"delete file.txt",
		"delete /dir",
		"delete ../dir",
	}

	for _, message := range tests {
		if _, err := NewCommand(message); err != nil {
			t.Errorf("Test Failed: \"%s\" inputted, an unexpected error occured: %s", message, err.Error())
		}
	}
}
