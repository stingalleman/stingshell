package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/stingalleman/stingshell/util"
)

// RunCmd execute command.
func RunCmd(input string) error {
	input = strings.TrimSuffix(input, "\n")
	homeDir, _ := os.UserHomeDir()

	input = strings.ReplaceAll(input, "~", homeDir)

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
			util.ExitShell()
		}
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
