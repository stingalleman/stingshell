package config

import (
	"os"
	"path"
)

// OpenFiles open files (and create if they do not exist)
func OpenFiles() (historyFile *os.File, configFile *os.File) {
	homeDir, _ := os.UserHomeDir()
	historyPath := path.Join(homeDir, "/.sting_history")
	configPath := path.Join(homeDir, "/.stingrc")

	if _, err := os.Stat(historyPath); os.IsNotExist(err) {
		os.Create(historyPath)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Create(configPath)
	}

	historyFile, _ = os.Open(historyPath)
	configFile, _ = os.Open(configPath)
	return historyFile, configFile
}

// CloseFiles yeah, just close the files
func CloseFiles(historyFile *os.File, configFile *os.File) {
	historyFile.Close()
	configFile.Close()
}

// WriteHistory write history
func WriteHistory(cmd string, historyFile *os.File) {
	historyFile.WriteString(cmd)
	historyFile.Sync()
}
