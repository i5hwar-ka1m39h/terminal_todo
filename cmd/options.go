package cmd

import (
	"os"
	"path/filepath"

	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
)

func GetAll() string {
	utils.Initiate()
	homedir, err := os.UserHomeDir()
	utils.CheckError(err)
	csvPath := filepath.Join(homedir, ".local", "share", "todoer", "todos.csv")
	data, err := os.ReadFile(csvPath)

	utils.CheckError(err)

	return string(data)
}
