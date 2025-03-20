package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimRight(command, "\n")
		args := strings.Split(command, " ")

		switch args[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
