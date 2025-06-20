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

	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{
		X: screenSize.X / 2,
		Y: 120,
	}, 16)

	buttonWidth := float32(480)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)
	startY := screenSize.Y / 2

	x := screenSize.X / 20
	y := startY

	mainMenu := &world.uistate.mainMenu

	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Start", mainMenu.selected == 1) {
		mainMenu.selected = 1
	} else if mainMenu.selected == 1 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Options", mainMenu.selected == 2) {
		mainMenu.selected = 2
	} else if mainMenu.selected == 2 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}

	x += buttonWidth * 1.1
	y = startY - buttonHeight/2
	if mainMenu.selected == 1 {
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
	} else if mainMenu.selected == 2 {
		renderOptions(world, rl.Vector2{X: x, Y: y})
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
