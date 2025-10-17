package cmd

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
)

func GetAll() [][]string {
	utils.Initiate()
	homedir, err := os.UserHomeDir()
	utils.CheckError(err)
	csvPath := filepath.Join(homedir, ".local", "share", "todoer", "todos.csv")
	file, err := os.Open(csvPath)
	utils.CheckError(err)

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	utils.CheckError(err)

	return records
}
