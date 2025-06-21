package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	selected uint8
}

func renderMainMenu(world *World) {
	screenSize := util.ScreenSize()
	rl.DrawRectangleV(rl.Vector2Zero(), screenSize, rl.ColorAlpha(rl.Black, 0.9))

	buttonWidth := float32(480)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)
	startY := buttonHeight + buttonSpacing

	x := screenSize.X / 20
	y := startY

	mainMenu := &world.uistate.mainMenu

	raygui.LabelButton(rl.Rectangle{X: screenSize.X - buttonWidth*1.5, Y: y, Width: buttonWidth * 1.5, Height: buttonHeight}, "alicevszombies")

	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Start", mainMenu.selected == 1) {
		mainMenu.selected = 1
	} else if mainMenu.selected == 1 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	raygui.Disable()
	raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Goals", mainMenu.selected == 2)
	raygui.Enable()

	y += buttonHeight + buttonSpacing
	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Stats", mainMenu.selected == 2) {
		mainMenu.selected = 2
	} else if mainMenu.selected == 2 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Options", mainMenu.selected == 3) {
		mainMenu.selected = 3
	} else if mainMenu.selected == 3 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}

	x += buttonWidth * 1.1
	y = startY - buttonHeight/2
	switch mainMenu.selected {
	case 1:
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Easy") {
			startGame(world, EASY)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Normal") {
			startGame(world, NORMAL)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Hard") {
			startGame(world, HARD)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Lunatic") {
			startGame(world, LUNATIC)
		}
	case 2:
		renderStats(rl.Vector2{X: x, Y: y})
	case 3:
		renderOptions(rl.Vector2{X: x, Y: y})
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
