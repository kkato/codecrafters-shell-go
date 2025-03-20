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
		case "type":
			switch args[1] {
			case "echo", "type", "exit":
				fmt.Println(args[1] + " is a shell builtin")
			default:
				pathEnv := os.Getenv("PATH")
				paths := strings.Split(pathEnv, ":")
				found := false
				for _, path := range paths {
					fullPath := path + "/" + args[1]
					if _, err := os.Stat(fullPath); err == nil {
						fmt.Println(args[1] + " is " + fullPath)
						found = true
						break
					}
				}
				if !found {
					fmt.Println(args[1] + ": not found")
				}
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
