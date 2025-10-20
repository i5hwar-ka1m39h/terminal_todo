package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

func AddTodo(todo, priority, timeRequired string) string {
	Initiate()
	homeDir, err := os.UserHomeDir()
	utils.CheckError(err)
	csvPath := filepath.Join(homeDir, ".local", "share", "todoer", "todos.csv")

	file, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, 0o644)
	utils.CheckError(err)
	defer file.Close()

	id := uuid.New()
	t := time.Now()
	create_at := t.Format("Monday, January 2, 2006 at 3:04 PM")
	status := "inProgress"

	writer := csv.NewWriter(file)
	defer writer.Flush()

	dataToAdd := []string{id.String(), todo, status, create_at, timeRequired, priority}

	writer.Write(dataToAdd)

	return id.String()
}

var addTodocmd = &cobra.Command{
	Use:   "add",
	Short: "Adds todo",
	Long:  "Adds todo with priority and time required",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		priority, _ := cmd.Flags().GetString("priority")
		hours, _ := cmd.Flags().GetInt("hours")

		res := AddTodo(title, priority, fmt.Sprintf("%d", hours))
		fmt.Printf("id : %s\n", res)
	},
}

func init() {
	addTodocmd.Flags().StringP("title", "t", "", "title of the todo")
	addTodocmd.Flags().StringP("priority", "p", "", "priority of the todo")
	addTodocmd.Flags().IntP("hours", "o", 0, "hours needed to complete the todo")
	rootCommand.AddCommand(addTodocmd)
	addTodocmd.MarkFlagRequired("title")
}
