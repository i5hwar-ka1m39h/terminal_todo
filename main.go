package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("reading a writing to file")

	//	data, err := os.ReadFile("./todo.csv")
	//	checkError(err)
	//	fmt.Println(string(data))
	//	d1 := []byte("shit i am writing checking whether it appends or not")
	//
	//	err := os.WriteFile("./todo.csv", d1, 0o644)
	//	checkError(err)
	//
	//	fileData, err := os.ReadFile("./todo.csv")
	//	checkError(err)
	//	fmt.Println(string(fileData))
	filePath := "."
	fileName := "todo.csv"
	fileExist := doesCSVExist(filePath, fileName)

	filePlace := filepath.Join(filePath, fileName)
	if !fileExist {
		// create file and write to it
		file, err := os.Create(filePlace)
		checkError(err)
		defer file.Close()

		todoString := "id, task, isDone, createAt, timeLimit, priority \n"
		_, err = file.WriteString(todoString)
	}

	file, err := os.OpenFile(filePlace, os.O_APPEND|os.O_WRONLY, 0o644)
	checkError(err)
	defer file.Close()

	todoString := "id_2, water the plants, false, 3 pm, 2hr, medium"
	_, err = file.WriteString(todoString)

	fileData, err := os.ReadFile(filePlace)
	checkError(err)

	fmt.Println(string(fileData))
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func doesCSVExist(dirName, fileName string) bool {
	dir, err := os.Open(dirName)
	checkError(err)

	allFileInfo, err := dir.ReadDir(-1)
	checkError(err)

	for _, file := range allFileInfo {
		if !file.IsDir() && file.Name() == fileName {
			return true
		}
	}
	return false
}
