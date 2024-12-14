package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")

		switch args[0] {
		case "exit":
			if args[1] == "0" {
				os.Exit(0)
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "type":
			if len(args) > 1 {
				switch args[1] {
				case "echo", "exit", "type":
					fmt.Printf("%s is a shell builtin\n", args[1])
				default:
					fmt.Printf("%s: not found\n", args[1])
				}
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
