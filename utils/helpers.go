package utils

import (
	"os"

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
			FG: renderer.Colors{color.FgWhite},
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
