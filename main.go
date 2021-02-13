package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	reader := bufio.NewReader(os.Stdin)
	homeDir, _ := os.UserHomeDir()

	for {
		currentDir, _ := os.Getwd()

		if currentDir == homeDir {
			currentDir = "~"
		}

		fmt.Printf("\n%s\n", currentDir)
		fmt.Printf("%s@%s >> ", user.Username, hostname)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		{
			if len(args) < 2 {
				homeDir, _ := os.UserHomeDir()

				return os.Chdir(homeDir)
			}
			return os.Chdir(args[1])
		}
	case "exit":
		{
			os.Exit(0)
		}
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
