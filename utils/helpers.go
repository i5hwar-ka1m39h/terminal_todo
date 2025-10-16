package utils

import (
	"os"
	"path/filepath"
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

		_, err = file.WriteString("id, task, status, created_at, time_limit, priority\n")
		CheckError(err)

	}
}
