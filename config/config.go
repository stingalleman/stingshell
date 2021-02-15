package config

import (
	"os"
	"path"
)

var (
	// Files config & history file
	Files filesStruct
)

// Files files
type filesStruct struct {
	historyFile os.File
	configFile  os.File
}

// AppendHistory writes input to the history file.
func (files *filesStruct) AppendHistory(input string) {
	files.historyFile.WriteString(input)
}

// OpenFiles open files (and create if they do not exist)
func (files *filesStruct) OpenFiles() {
	homeDir, _ := os.UserHomeDir()
	historyPath := path.Join(homeDir, "/.sting_history")
	configPath := path.Join(homeDir, "/.stingrc")

	if _, err := os.Stat(historyPath); os.IsNotExist(err) {
		os.Create(historyPath)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Create(configPath)
	}

	historyFile, _ := os.OpenFile(historyPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	configFile, _ := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	files.configFile = *configFile
	files.historyFile = *historyFile
}

// CloseFiles yeah, just close the files
func (files *filesStruct) CloseFiles() {
	files.historyFile.Close()
	files.configFile.Close()

}
