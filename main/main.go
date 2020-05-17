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
	var grid = createGrid(int(difficulty))
	//var grid [9][9]cell
	/*grid = [9][9] cell{
		{{3, true}, {0, false}, {6, true}, {5, true}, {0, false}, {8, true}, {4, true}, {0, false}, {0, false}},
		{{5, true}, {2, true}, {0, false}, {0, false}, {0, false}, {0, false}, {0, false}, {0, false}, {0, false}},
		{{0, false}, {8, true}, {7, true}, {0, false}, {0, false}, {0, false}, {0, false}, {3, true}, {1, true}},
		{{0, false}, {0, false}, {3, true}, {0, false}, {1, true}, {0, false}, {0, false}, {8, true}, {0, false}},
		{{9, true}, {0, false}, {0, false}, {8, true}, {6, true}, {3, true}, {0, false}, {0, false}, {5, true}},
		{{0, false}, {5, true}, {0, false}, {0, false}, {9, true}, {0, false}, {6, true}, {0, false}, {0, false}},
		{{1, true}, {3, true}, {0, false}, {0, false}, {0, false}, {0, false}, {2, true}, {5, true}, {0, false}},
		{{0, false}, {0, false}, {0, false}, {0, false}, {0, false}, {0, false}, {0, false}, {7, true}, {4, true}},
		{{0, false}, {0, false}, {5, true}, {2, true}, {0, false}, {6, true}, {3, true}, {0, false}, {0, false}},
	}*/
	displayGrid(&grid)

	solveGrid(&grid)
	displayGrid(&grid)
}
