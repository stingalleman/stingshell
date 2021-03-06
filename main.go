package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"runtime"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/stingalleman/stingshell/cmd"
	"github.com/stingalleman/stingshell/config"
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

	// windows circlejerk
	if runtime.GOOS == "windows" {
		fgWhite := color.New(color.FgWhite).SprintFunc()
		red := color.New(color.BgRed).SprintFunc()
		fmt.Fprint(color.Output, bold(fgWhite(red("\n\n\ngast, ben je serieus fucking windows aan het gebruiken? ga fucking linux of macos gebruiken ofzo. fucking borderline cretin, vieze gremlin creature\n\nmongool.\n\n"))))
	}

	config.OpenFiles()
	setupCloseHandler()

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
		config.HistoryFile.AppendHistory(input)

		if err = cmd.Exec(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		util.ExitShell()
	}()
}
