package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

func UpdateTodo(id string) (string, error) {
	homeDir, err := os.UserHomeDir()
	utils.CheckError(err)
	csvPath := filepath.Join(homeDir, ".local", "share", "todoer", "todos.csv")

	file, err := os.Open(csvPath)
	utils.CheckError(err)

	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	utils.CheckError(err)

	found := false
	for i, row := range record {
		if i == 0 {
			continue
		}
		if row[0] == id {
			row[2] = "done"
			found = true
			break
		}
	}

	if !found {
		return "", errors.New("the todo with given id not found")
	}

	file, err = os.Create(csvPath)
	utils.CheckError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(record)
	return "todo marked as done", err
}

var markDoneCmd = &cobra.Command{
	Use:   "markDone",
	Short: "mark the todo as done",
	Long:  "change the status of the todo with given id as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := UpdateTodo(args[0])
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s\n", result)
	},
}

func init() {
	rootCommand.AddCommand(markDoneCmd)
}
