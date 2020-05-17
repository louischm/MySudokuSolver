package main

func solveGrid(grid *[9][9]cell) {
	var result bool
	var x, y, value int

	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			if !grid[x][y].original {
				grid[x][y].value++
				result, value = solveCell(grid, &x, &y, &grid[x][y].value)
				grid[x][y].value = value
				//displayGrid(grid)
				if !result {
					x, y = backtrack(grid, &x, &y)
					if y == 9 {
						y = -1
						x++
					} else {
						y--
					}
				} else {
					if x == 8 && y == 8 {
						displayGrid(grid)
						return
					}
				}
			} else {
				if x == 8 && y == 8 {
					displayGrid(grid)
					return
				}
				if y == 9 {
					y = -1
					x++
				}
			}
		}
	}
	displayGrid(grid)
}

func backtrack(grid *[9][9]cell, posX *int, posY *int) (int, int) {
	var x, y int

	for x = *posX; x >= 0; x-- {
		for y = *posY; y >= 0; y-- {

			if x < 0 {
				x = 0
			}

			if (x != *posX || y != *posY) && !grid[x][y].original {
				if x == 0 && y == 0 {
					return x, y
				}
				return x, y
			}

			if y == 0 {
				y = 9
				x--
			}
		}
	}
	return 0, 0
}

func solveCell(grid *[9][9]cell, x *int, y *int, value *int) (bool, int) {
	if *value > 9 {
		return false, 0
	}

	if !checkCell(x, y, value, grid) {
		*value++
		return solveCell(grid, x, y, value)
	} else {
		return true, *value
	}
}
