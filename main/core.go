package main

import (
	"github.com/oakmound/oak"
	"log"
)

func run() {
	err := oak.Add("menu", MenuStart, MenuLoop, MenuEnd)
	if err != nil {
		log.Fatalln(err)
	}

	err = oak.Add("sudoku", SudokuStart, SudokuLoop, SudokuEnd)
	if err != nil {
		log.Fatalln(err)
	}

	oak.Init("menu")
}
