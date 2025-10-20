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

func DeleteTodo(id string) (string, error) {
	homeDir, err := os.UserHomeDir()
	utils.CheckError(err)

	csvPath := filepath.Join(homeDir, ".local", "share", "todoer", "todos.csv")

	file, err := os.Open(csvPath)
	utils.CheckError(err)

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	utils.CheckError(err)

	newRecords := [][]string{}
	found := false

	for i, row := range records {
		if i == 0 {
			newRecords = append(newRecords, row)
			continue
		}

		if row[0] == id {
			found = true
			break
		}

		newRecords = append(newRecords, row)
	}

	if !found {
		return "", errors.New("todo with given id not found")
	}

	tempFile := csvPath + ".temp"

	tmp, err := os.Create(tempFile)
	utils.CheckError(err)

	writer := csv.NewWriter(tmp)
	if err := writer.WriteAll(newRecords); err != nil {
		tmp.Close()
		return "", err
	}

	writer.Flush()
	tmp.Close()

	if err := os.Rename(tempFile, csvPath); err != nil {
		return "", err
	}

	return "todo deleted successfully", nil
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete todo",
	Long:  "remove the todo with given id from systme",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please provide the id of the todo")
		} else {
			res, err := DeleteTodo(args[0])
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("%s\n", res)
		}
	},
}

func init() {
	rootCommand.AddCommand(deleteCmd)
}
