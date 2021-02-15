package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/fatih/color"
	"github.com/stingalleman/stingshell/util"
)

func main() {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	reader := bufio.NewReader(os.Stdin)
	homeDir, _ := os.UserHomeDir()

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	for {
		currentDir, _ := os.Getwd()

		if currentDir == homeDir {
			currentDir = "~"
		} else if strings.Contains(currentDir, homeDir) {
			currentDir = "~" + strings.SplitAfter(currentDir, homeDir)[1]
		}

		fmt.Printf("\n%s\n", yellow(currentDir))
		fmt.Printf("%s%s%s %s ", green(user.Username), green("@"), green(hostname), bold(green("❯")))

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = util.RunCmd(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
