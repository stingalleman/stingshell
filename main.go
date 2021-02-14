package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/stingalleman/stingshell/util"
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
		} else if strings.Contains(currentDir, homeDir) {
			currentDir = "~" + strings.SplitAfter(currentDir, homeDir)[1]
		}

		fmt.Printf("\n%s\n", currentDir)
		fmt.Printf("%s@%s >> ", user.Username, hostname)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = util.RunCmd(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
