package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	cmd "github.com/jushutch/dfs/internal/command"
)

// Command line client used to interact with the
// distributed file system
type Client struct {
	outPort string
}

// Continuously recieve and interpret user input
func (client Client) Interpret() error {
	for {
		fmt.Print(string("\033[31m"), "root@client", string("\033[37m"), ": ")
		in := bufio.NewReader(os.Stdin)
		command, err := in.ReadString('\n')
		if err != nil {
			return err
		}
		fsCommand, err := cmd.NewCommand(strings.TrimSpace(command))
		if err != nil {
			fmt.Printf("\t%v\n", err)
			continue
		}
		err = client.Send(fsCommand.String())
		if err != nil {
			return err
		}
	}
}

func (client Client) Send(msg string) error {
	conn, err := net.Dial("tcp", ":"+client.outPort)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = conn.(*net.TCPConn).CloseWrite()
	if err != nil {
		return err
	}
	return nil
}
