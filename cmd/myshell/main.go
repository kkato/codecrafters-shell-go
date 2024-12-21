package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			if len(args) > 1 && args[1] == "0" {
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
			path, found := findExecutable(args[0])
			if found {
				runExternalCommand(path, args[1:])
			} else {
				fmt.Printf("%s: command not found\n", args[0])
			}
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

func runExternalCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %v\n", err)
	}
}
