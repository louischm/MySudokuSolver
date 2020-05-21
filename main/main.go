package main

import (
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args[1:])

	if argLength > 1 {
		fmt.Println("This program only need one or zero argument: ./app [sudoku-file.txt]")
		return
	}

	result, err := verifyFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if argLength == 1 && !result {
		fmt.Println("Sudoku file wrong format, see sudoku-example.txt for more details")
		return
	}

	run()
}
