package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCmd: execute command.
func RunCmd(input string) error {
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
			fmt.Fprint(os.Stderr, "bye bye!\n")
			os.Exit(0)
		}
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
