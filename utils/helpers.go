package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

//	func DoesCSVExist(dirName, fileName string) bool {
//		dir, err := os.Open(dirName)
//		CheckError(err)
//
//		allFileInfo, err := dir.ReadDir(-1)
//		CheckError(err)
//
//		for _, file := range allFileInfo {
//			if !file.IsDir() && file.Name() == fileName {
//				return true
//			}
//		}
//		return false
//	}
func Initiate() {
	homeDir, err := os.UserHomeDir()
	CheckError(err)

	dataDir := filepath.Join(homeDir, ".local", "share", "todoer")

	err = os.MkdirAll(dataDir, 0o755)
	CheckError(err)

	// check if csv exist if not then create else exit
	csvPath := filepath.Join(dataDir, "todos.csv")
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		file, err := os.Create(csvPath)
		CheckError(err)

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		header := []string{"id", "title", "status", "created_at", "time_limit", "priority"}
		err = writer.Write(header)
		CheckError(err)

	}
}

func Tabelize(data [][]string) {
	if len(data) == 0 {
		color.Red("No records found")
		return
	}

	colorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgCyan, color.Bold},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgGreen},
		},
	}
	table := tablewriter.NewTable(
		os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
	)

	header := data[0]

	table.Header(header)

	table.Bulk(data[1:])

	table.Render()
}
