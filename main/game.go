package main

import (
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/scene"
	"image"
	"image/color"
	"log"
	"os"
	"sync"
	"time"
)

var valueFontGenerator = render.FontGenerator{
	File:    "luxisbi.ttf",
	Color:   image.NewUniform(color.RGBA{0, 0, 0, 255}),
	Size:    50,
	Hinting: "",
	DPI:     20,
}

type GridCounter struct {
	grid [9][9]cell
	mux  sync.Mutex
}

var once bool
var gridCounter GridCounter

func SudokuStart(prevScene string, inData interface{}) {
	once = false
	_, err := render.Draw(render.NewColorBox(1250, 1000, color.RGBA{255, 255, 255, 255}), 0)
	if err != nil {
		log.Fatalln(err)
	}

	renderVertLine()
	renderHorizonLine()

	gridCounter.mux = sync.Mutex{}

	if len(os.Args[1:]) == 1 {
		gridCounter.grid = loadGridFromFile(os.Args[1])
	} else {
		gridCounter.grid = createGrid()
	}

	gridCounter.mux.Lock()
	renderGrid(&gridCounter.grid)
	gridCounter.mux.Unlock()
}

func SudokuLoop() bool {
	if oak.IsDown("Escape") {
		oak.Quit()
	}

	result, timeHeld := oak.IsHeld("Spacebar")
	if result && timeHeld.Milliseconds() > 75 && timeHeld.Milliseconds() < 100 {
		if !once {
			go solveGrid(&gridCounter)
			once = true
		}
	}
	return true
}

func SudokuEnd() (nextScene string, result *scene.Result) {
	return "sudoku", nil
}

func renderGrid(grid *[9][9]cell) {
	var x, y int
	displayGrid(&gridCounter.grid)

	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {

			var i, j int
			yFont := 50

			for i = 0; i < 9; i++ {
				xFont := 125
				for j = 0; j < 9; j++ {
					if x == i && y == j {
						newCell(grid[x][y].value, float64(xFont), float64(yFont), x, y)
					}
					xFont += 50
				}
				yFont += 50
			}

		}
	}
}

func newCell(value int, x, y float64, realX, realY int) {
	if value > 0 {
		valueFontGenerator.Color = image.NewUniform(color.RGBA{0, 0, 0, 255})
	} else {
		valueFontGenerator.Color = image.NewUniform(color.RGBA{255, 255, 255, 255})
	}
	valueFont := valueFontGenerator.Generate()
	cellRender, err := render.Draw(valueFont.NewIntText(&value, x, y), 3)
	if err != nil {
		log.Fatalln(err)
	}
	cellEntity := entities.NewMoving(x, y, 25, 25, cellRender, nil, 0, 0)
	cellEntity.Bind(updateCell(realX, realY, &value, x, y), event.Enter)
}

func updateCell(x, y int, newValue *int, xFont, yFont float64) func(int, interface{}) int {
	return func(id int, nothing interface{}) int {
		if gridCounter.grid[x][y].value != *newValue {
			gridCounter.mux.Lock()
			p := event.GetEntity(id).(*entities.Moving)
			p.R.Undraw()

			if gridCounter.grid[x][y].value > 0 {
				valueFontGenerator.Color = image.NewUniform(color.RGBA{0, 200, 0, 255})
			} else {
				valueFontGenerator.Color = image.NewUniform(color.RGBA{200, 0, 0, 255})
			}
			valueFont := valueFontGenerator.Generate()
			valueRender, err := render.Draw(valueFont.NewIntText(&gridCounter.grid[x][y].value, xFont, yFont), id)
			if err != nil {
				log.Fatalln(err)
			}
			p.SetRenderable(valueRender)
			*newValue = gridCounter.grid[x][y].value
			gridCounter.mux.Unlock()
			time.Sleep(5 * time.Millisecond)
		}
		return 0
	}
}

func renderVertLine() {
	_, err := render.Draw(render.NewLine(100, 25, 100, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(101, 25, 101, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(150, 25, 150, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(200, 25, 200, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(250, 25, 250, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(251, 25, 251, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(300, 25, 300, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(350, 25, 350, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(400, 25, 400, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(401, 25, 401, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(450, 25, 450, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(500, 25, 500, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(550, 25, 550, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(551, 25, 551, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

}

func renderHorizonLine() {
	_, err := render.Draw(render.NewLine(100, 25, 551, 25, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(100, 26, 551, 26, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 75, 551, 75, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 125, 551, 125, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 175, 551, 175, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(100, 176, 551, 176, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 225, 551, 225, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 275, 551, 275, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 325, 551, 325, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(100, 326, 551, 326, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 375, 551, 375, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 425, 551, 425, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewLine(100, 475, 551, 475, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = render.Draw(render.NewLine(100, 476, 551, 476, color.RGBA{0, 0, 0, 255}), 1)
	if err != nil {
		log.Fatalln(err)
	}

}
