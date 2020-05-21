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
)

var titleFontGenerator = render.FontGenerator{
	File:    "luxisbi.ttf",
	Color:   image.NewUniform(color.RGBA{0, 0, 0, 255}),
	Size:    100,
	Hinting: "",
	DPI:     20,
}

var menuFontGenerator = render.FontGenerator{
	File:    "luxisbi.ttf",
	Size:    50,
	Hinting: "",
	DPI:     20,
}

var menuSelected = 1

func MenuStart(prevScene string, data interface{}) {
	err := oak.MoveWindow(350, 500, 1250, 1000)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = render.Draw(render.NewColorBox(1250, 1000, color.RGBA{255, 255, 255, 255}), 0)
	titleFont := titleFontGenerator.Generate()
	_, err = render.Draw(titleFont.NewStrText("Sudoku", 270, 50), 1)
	if err != nil {
		log.Fatalln(err)
	}
	newMenuBar(310, 165, "Play", 1)
	newMenuBar(310, 195, "Exit", 0)
}

func MenuLoop() bool {
	if oak.IsDown("Escape") {
		oak.Quit()
	}

	if oak.IsDown("ReturnEnter") {
		log.Println("Enter")

		if menuSelected == 0 {
			oak.Quit()
		} else if menuSelected == 1 {
			return false
		}
		return false
	}

	return true
}

func MenuEnd() (nextScene string, result *scene.Result) {
	return "sudoku", nil
}

func newMenuBar(x, y float64, title string, menuNumber int) {

	if menuSelected == menuNumber {
		menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 200, 0, 255})
	} else {
		menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 0, 0, 255})
	}

	menuFont := menuFontGenerator.Generate()

	menuRender, err := render.Draw(menuFont.NewStrText(title, x, y), 0)
	if err != nil {
		log.Fatalln(err)
	}

	menuEntity := entities.NewMoving(x, y, 100, 20, menuRender, nil, 0, 0)

	menuEntity.Bind(enterMenu("UpArrow", "DownArrow", menuNumber), event.Enter)
}

func enterMenu(up string, down string, menuNumber int) func(int, interface{}) int {
	return func(id int, nothing interface{}) int {
		p := event.GetEntity(id).(*entities.Moving)
		var title string
		var x, y float64

		if menuNumber == 0 {
			title = "Exit"
			x = 310
			y = 195
		}

		if menuNumber == 1 {
			title = "Play"
			x = 310
			y = 165
		}

		result, time := oak.IsHeld(up)
		if result && time.Milliseconds() > 75 && time.Milliseconds() < 100 {
			p.R.Undraw()

			if menuSelected < 1 {
				menuSelected++
			}

			if menuSelected == menuNumber {
				menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 200, 0, 255})
			} else {
				menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 0, 0, 255})
			}

			menuFont := menuFontGenerator.Generate()
			menuRender, err := render.Draw(menuFont.NewStrText(title, x, y), 0)
			if err != nil {
				log.Fatalln(err)
			}

			p.SetRenderable(menuRender)
		}

		result, time = oak.IsHeld(down)
		if result && time.Milliseconds() > 75 && time.Milliseconds() < 100 {
			p.R.Undraw()

			if menuSelected > 0 {
				menuSelected--
			}

			if menuSelected == menuNumber {
				menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 200, 0, 255})
			} else {
				menuFontGenerator.Color = image.NewUniform(color.RGBA{0, 0, 0, 255})
			}

			menuFont := menuFontGenerator.Generate()
			menuRender, err := render.Draw(menuFont.NewStrText(title, x, y), 0)
			if err != nil {
				log.Fatalln(err)
			}

			p.SetRenderable(menuRender)
		}

		return 0
	}
}
