package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type cell struct {
	value    int
	original bool
}

func loadGridFromFile(filename string) [9][9]cell {
	var grid [9][9]cell
	var i int
	var char string
	x := 0
	y := 0

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	for i = 0; i < len(data); i++ {
		if data[i] == 0 || x == 9 {
			break
		} else if data[i] >= '0' && data[i] <= '9' {
			char = string(data[i])
			log.Printf("x: %d, y: %d\n", x, y)
			grid[x][y].value, err = strconv.Atoi(char)
			if err != nil {
				log.Fatalln(err)
			}

			if grid[x][y].value != 0 {
				grid[x][y].original = true
			}
			y++

		} else if data[i] == '\n' {
			x++
			y = 0
		}
	}

	displayGrid(&grid)
	return grid
}

func createGrid() [9][9]cell {
	fmt.Println("Creating sudoku grid")
	var grid [9][9]cell
	var x, y int

	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			newCell := generateCell()
			for true {
				if !checkCell(&x, &y, &newCell.value, &grid) {
					newCell = reGenerateCell()
				} else {
					break
				}
			}
			grid[x][y] = newCell
		}
	}
	return grid
}

func generateCell() cell {
	rand.Seed(time.Now().UnixNano())
	var newCell cell

	minEmpty := 0
	maxEmpty := 4
	resultEmpty := rand.Intn(maxEmpty-minEmpty+1) + minEmpty

	if resultEmpty != 4 {
		newCell = cell{0, false}
	} else {
		minValue := 1
		maxValue := 9
		newCell = cell{rand.Intn(maxValue-minValue+1) + minValue, true}
	}
	return newCell
}

func reGenerateCell() cell {
	rand.Seed(time.Now().UnixNano())
	var newCell cell

	minValue := 1
	maxValue := 9
	newCell = cell{rand.Intn(maxValue-minValue+1) + minValue, true}
	return newCell
}

func checkCell(xNewCell *int, yNewCell *int, newValue *int, grid *[9][9]cell) bool {
	if *newValue == 0 {
		return true
	}

	if *newValue > 9 {
		return false
	}

	if !checkColumn(yNewCell, xNewCell, newValue, grid) {
		return false
	}

	if !checkLine(xNewCell, yNewCell, newValue, grid) {
		return false
	}

	if !checkSubGrid(xNewCell, yNewCell, newValue, grid) {
		return false
	}
	return true
}

func checkColumn(yNewCell *int, xNewCell *int, newValue *int, grid *[9][9]cell) bool {
	var x int

	for x = 0; x < 9; x++ {
		if *newValue == grid[x][*yNewCell].value && *xNewCell != x {
			return false
		}
	}

	return true
}

func checkLine(xNewCell *int, yNewCell *int, newValue *int, grid *[9][9]cell) bool {
	var y int

	for y = 0; y < 9; y++ {
		if *newValue == grid[*xNewCell][y].value && *yNewCell != y {
			return false
		}
	}
	return true
}

func checkSubGrid(xNewCell *int, yNewCell *int, newValue *int, grid *[9][9]cell) bool {
	var xSubGrid, ySubGrid, x, y int

	for x = 3; x < 10; x += 3 {
		if *xNewCell/x == 0 {
			xSubGrid = x - 3
			break
		}
	}

	for y = 3; y < 10; y += 3 {
		if *yNewCell/y == 0 {
			ySubGrid = y - 3
			break
		}
	}

	for x = xSubGrid; x < xSubGrid+3; x++ {
		for y = ySubGrid; y < ySubGrid+3; y++ {
			if *newValue == grid[x][y].value && (x != *xNewCell && y != *yNewCell) {
				return false
			}
		}
	}
	return true
}

func displayGrid(grid *[9][9]cell) {
	var x, y int

	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			fmt.Printf("%d", grid[x][y].value)
			if y == 2 || y == 5 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
		if x == 2 || x == 5 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n\n\n")
}

func verifyFile(fileName string) (bool, error) {
	var i, num, retLine int
	var data []byte

	if !strings.HasSuffix(fileName, ".txt") {
		return false, nil
	}

	_, err := os.Open(fileName)
	if err != nil {
		return false, err
	}

	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		return false, err
	}

	for i = 0; i < len(data); i++ {
		if data[i] < '0' || data[i] > '9' {
			if data[i] != '\n' && data[i] != 0 && data[i] != ' ' {
				return false, nil
			} else if data[i] == '\n' {
				retLine++
			}
		} else {
			num++
		}
	}

	if num != 81 || retLine <= 7 {
		return false, nil
	}

	return true, nil
}
