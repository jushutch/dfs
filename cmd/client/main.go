package main

import (
	"fmt"
	"os"
)

func main() {
	client := Client{os.Args[1]}
	err := client.Interpret()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
