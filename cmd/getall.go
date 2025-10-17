package cmd

import (
	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

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
