package util

import (
	"fmt"
	"os"

	"github.com/stingalleman/stingshell/config"
)

// ExitShell exit shell
func ExitShell() {
	config.Files.CloseFiles()
	fmt.Fprint(os.Stderr, "\nbye bye!\n")

	os.Exit(0)
}
