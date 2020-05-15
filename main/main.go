package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	argLength := len(os.Args[1:])

	if argLength != 1 {
		fmt.Println("This program only need one argument: ./app [1-5]")
		return
	}

	difficulty, err := strconv.ParseInt(os.Args[1], 0, 64)

	if err != nil {
		log.Fatal(err)
		return
	}

	if difficulty < 1 || difficulty > 5 {
		fmt.Println("Please provide a number between 1 to 5 for difficulty")
	}
	createGrid(int(difficulty))
}
