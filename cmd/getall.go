package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAllcmd = &cobra.Command{
	Use:   "getAll",
	Short: "get all todo",
	Long:  "get all the todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", GetAll())
	},
}

func init() {
	rootCommand.AddCommand(getAllcmd)
}
