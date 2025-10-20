package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/i5hwar-ka1m39h/terminal_todo/utils"
	"github.com/spf13/cobra"
)

const asciiBanner = `
 ________               __                               
/        |             /  |                              
$$$$$$$$/______    ____$$ |  ______    ______    ______  
   $$ | /      \  /    $$ | /      \  /      \  /      \ 
   $$ |/$$$$$$  |/$$$$$$$ |/$$$$$$  |/$$$$$$  |/$$$$$$  |
   $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |$$    $$ |$$ |  $$/ 
   $$ |$$ \__$$ |$$ \__$$ |$$ \__$$ |$$$$$$$$/ $$ |      
   $$ |$$    $$/ $$    $$ |$$    $$/ $$       |$$ |      
   $$/  $$$$$$/   $$$$$$$/  $$$$$$/   $$$$$$$/ $$/       
                                                         
                                                         
                                                         
A terminal todo companion.
	`

func Initiate() {
	homeDir, err := os.UserHomeDir()
	utils.CheckError(err)

	dataDir := filepath.Join(homeDir, ".local", "share", "todoer")

	err = os.MkdirAll(dataDir, 0o755)
	utils.CheckError(err)

	// check if csv exist if not then create else exit
	csvPath := filepath.Join(dataDir, "todos.csv")
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		file, err := os.Create(csvPath)
		utils.CheckError(err)

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		header := []string{"id", "title", "status", "created_at", "time_limit", "priority"}
		err = writer.Write(header)
		utils.CheckError(err)

	}
}

var rootCommand = &cobra.Command{
	Use:   "todoer",
	Short: "a terminal based task tracker",
	Long:  "Maintain you productivity inside the termial with single command.",
	Run: func(cmd *cobra.Command, args []string) {
		Initiate()
		color.Cyan(asciiBanner)

		color.Magenta("Available Commands:")
		cmdColor := color.New(color.FgRed).SprintFunc()
		descColor := color.New(color.FgGreen).SprintFunc()

		for _, c := range cmd.Commands() {
			if !c.Hidden {
				fmt.Printf(" %s -> %s\n", cmdColor(fmt.Sprintf("%-10s", c.Use)), descColor(c.Short))
			}
		}
		color.Magenta("\nRun 'todoer help <Command>' for detailed description for the command.")
	},
}

func Executer() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing the the commands: %s\n", err)
		os.Exit(1)
	}
}
