package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/stingalleman/stingshell/util"
)

// Exec executes input as command.
func Exec(input string) error {
	input = strings.TrimSuffix(input, "\n")
	homeDir, _ := os.UserHomeDir()

	input = strings.ReplaceAll(input, "~", homeDir)
	input = os.ExpandEnv(input)

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
	case "export":
		{
			if len(args) <= 1 {
				fmt.Print(os.Environ())
			}
			for i := 1; i < len(args); i++ {
				values := strings.Split(args[i], "=")
				os.Setenv(values[0], values[1])
			}
			return errors.New("")
		}
	case "exit":
		{
			util.ExitShell()
		}
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
