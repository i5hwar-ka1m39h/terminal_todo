package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

func AddTodo(todo, priority, timeRequired string) string {
	utils.Initiate()
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

	entryString := fmt.Sprintf(
		"%s, %s, %s, %s, %s, %s \n",
		id.String(),
		todo,
		status,
		create_at,
		timeRequired,
		priority,
	)
	_, err = file.WriteString(entryString)

	return entryString
}

var addTodocmd = &cobra.Command{
	Use:   "add",
	Short: "Adds todo",
	Long:  "Adds todo with priority and time required",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			res := AddTodo(args[0], args[1], args[2])
			fmt.Println(res)
		} else {
			fmt.Println("please provide three arguments title, priority, time required")
		}
	},
}

func init() {
	rootCommand.AddCommand(addTodocmd)
}
