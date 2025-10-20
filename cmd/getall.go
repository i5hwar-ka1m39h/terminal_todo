package cmd

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

func GetAll() [][]string {
	Initiate()
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

var getAllcmd = &cobra.Command{
	Use:   "getAll",
	Short: "get all todo",
	Long:  "get all the todo",
	Run: func(cmd *cobra.Command, args []string) {
		res := GetAll()
		utils.Tabelize(res)
	},
}

func init() {
	rootCommand.AddCommand(getAllcmd)
}
