package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	client := Client{outPort: port}
	err = client.Interpret()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
