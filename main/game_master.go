package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cell struct {
	value    int
	original bool
}

func createGrid(difficulty int) {
	fmt.Printf("Creating grid with difficulty %d\n", difficulty)
	var grid [9][9]cell
	var x, y int

	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			newCell := generateCell(difficulty)
			for true {
				if !checkCell(x, y, newCell.value, grid) {
					newCell = reGenerateCell()
				} else {
					break
				}
			}
			grid[x][y] = newCell
		}
	}
	displayGrid(grid)
}

func generateCell(difficulty int) cell {
	rand.Seed(time.Now().UnixNano())
	var newCell cell

	minEmpty := 0
	maxEmpty := difficulty
	resultEmpty := rand.Intn(maxEmpty-minEmpty+1) + minEmpty

	if resultEmpty != difficulty {
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

func checkCell(xNewCell int, yNewCell int, newValue int, grid [9][9]cell) bool {
	if newValue == 0 {
		return true
	}

	if !checkColumn(yNewCell, newValue, grid) {
		return false
	}

	if !checkLine(xNewCell, newValue, grid) {
		return false
	}

	if !checkSubGrid(xNewCell, yNewCell, newValue, grid) {
		return false
	}
	return true
}

func checkColumn(yNewCell int, newValue int, grid [9][9]cell) bool {
	var x int

	for x = 0; x < 9; x++ {
		if newValue == grid[x][yNewCell].value {
			return false
		}
	}

	return true
}

func checkLine(xNewCell int, newValue int, grid [9][9]cell) bool {
	var y int

	for y = 0; y < 9; y++ {
		if newValue == grid[xNewCell][y].value {
			return false
		}
	}
	return true
}

func checkSubGrid(xNewCell int, yNewCell int, newValue int, grid [9][9]cell) bool {
	var xSubGrid, ySubGrid, x, y int

	for x = 3; x < 10; x += 3 {
		if xNewCell/x == 0 {
			xSubGrid = x - 3
			break
		}
	}

	for y = 3; y < 10; y += 3 {
		if yNewCell/y == 0 {
			ySubGrid = y - 3
			break
		}
	}

	for x = xSubGrid; x < xSubGrid+3; x++ {
		for y = ySubGrid; y < ySubGrid+3; y++ {
			if newValue == grid[x][y].value {
				return false
			}
		}
	}
	return true
}

func displayGrid(grid [9][9]cell) {
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
}
