package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
					path, found := findExecutable(args[1])
					if found {
						fmt.Println(path)
					} else {
						fmt.Printf("%s: not found\n", args[1])
					}
				}
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

func findExecutable(command string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return "", false
	}

	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	for _, dir := range paths {
		fullPath := filepath.Join(dir, command)
		if fileInfo, err := os.Stat(fullPath); err == nil {
			if !fileInfo.IsDir() && (fileInfo.Mode()&0111 != 0) {
				return fullPath, true
			}
		}
	}

	return "", false
}
