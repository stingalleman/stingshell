package config

import (
	"os"
	"path"
)

var (
	// HistoryFile history file (sting_history)
	HistoryFile historyFileStruct

	// ConfigFile config file (stingrc)
	ConfigFile configFileStruct
)

// Files files
type historyFileStruct struct {
	os.File
}

type configFileStruct struct {
	os.File
}

// AppendHistory writes input to the history file.
func (h *historyFileStruct) AppendHistory(input string) {
	h.WriteString(input)
}

// OpenFiles open files (and create if they do not exist)
func (h *historyFileStruct) openHistoryFile() {
	homeDir, _ := os.UserHomeDir()
	historyPath := path.Join(homeDir, "/.sting_history")

	if _, err := os.Stat(historyPath); os.IsNotExist(err) {
		os.Create(historyPath)
	}

	historyFile, _ := os.OpenFile(historyPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	h.File = *historyFile
}

func (c *configFileStruct) openConfigFile() {
	homeDir, _ := os.UserHomeDir()
	configPath := path.Join(homeDir, "/.stingrc")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Create(configPath)
	}

	configFile, _ := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	c.File = *configFile
}

func (c *configFileStruct) closeConfigFile() {
	c.Close()
}

func (h *historyFileStruct) closeHistoryFile() {
	h.Close()
}

// OpenFiles open config & history file
func OpenFiles() {
	ConfigFile.openConfigFile()
	HistoryFile.openHistoryFile()
}

// CloseFiles yeah, just close the files
func CloseFiles() {
	ConfigFile.closeConfigFile()
	HistoryFile.closeHistoryFile()
}
