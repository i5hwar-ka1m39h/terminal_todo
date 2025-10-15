package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("reading a writing to file")

	data, err := os.ReadFile("./todo.csv")
	checkError(err)
	fmt.Println(string(data))
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
