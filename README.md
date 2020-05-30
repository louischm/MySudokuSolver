# MySudokuSolver
Sudoku solver written in Go using backtracking algorithm.

## Prerequisite

In order to run this project you'll need:

* Go 1.14 or higher
* Install Oak golang package: https://github.com/oakmound/oak

## Run

You can now run the program like so:

`go run main/`

Or if you want to load one of your own sudoku grid, do it like so :

`go run main/ [your file].txt`

You can use the example that are in the `main/` directory

## Usage

In the 'Menu' screen you can :

* Choose between two options:
    * Play: Generate a sudoku grid if you haven't give one to the program
    * Exit: Exit the program
* Exit the program by pressing the Escape button
* Choose an option by pressing the Enter button

In the 'Play' screen you can start and see the solving algorithm by pressing the Space button.