package main

import (
	"github.com/i5hwar-ka1m39h/terminal_todo/cmd"
)

func main() {
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
	//	filePath := "."
	//	fileName := "todo.csv"
	//	fileExist := doesCSVExist(filePath, fileName)
	//
	//	filePlace := filepath.Join(filePath, fileName)
	//	if !fileExist {
	//		// create file and write to it
	//		file, err := os.Create(filePlace)
	//		checkError(err)
	//		defer file.Close()
	//
	//		todoString := "id, task, isDone, createAt, timeLimit, priority \n"
	//		_, err = file.WriteString(todoString)
	//	}
	//
	//	file, err := os.OpenFile(filePlace, os.O_APPEND|os.O_WRONLY, 0o644)
	//	checkError(err)
	//	defer file.Close()
	//
	//	todoString := "id_2, water the plants, false, 3 pm, 2hr, medium\n"
	//	_, err = file.WriteString(todoString)
	//
	//	fileData, err := os.ReadFile(filePlace)
	//	checkError(err)
	//
	//	fmt.Println(string(fileData))
	cmd.Executer()
}
