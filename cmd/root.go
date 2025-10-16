package cmd

import (
	"fmt"
	"os"

	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "todoer",
	Short: "a terminal based task tracker",
	Long:  "Maintain you productivity inside the termial with single command.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Initiate()
	},
}

func Executer() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing the the commands: %s\n", err)
		os.Exit(1)
	}
}
