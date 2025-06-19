package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	difficultySelect bool
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

	if world.uistate.mainMenu.difficultySelect {
		raygui.SetState(raygui.STATE_PRESSED)
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Start") {
			world.uistate.mainMenu.difficultySelect = !world.uistate.mainMenu.difficultySelect
		}
		raygui.SetState(raygui.STATE_NORMAL)
	} else {
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Start") {
			world.uistate.mainMenu.difficultySelect = !world.uistate.mainMenu.difficultySelect
		}
	}

	y += buttonHeight + buttonSpacing

	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen") {
		rl.ToggleFullscreen()
	}

	y += buttonHeight + buttonSpacing

	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}

	if world.uistate.mainMenu.difficultySelect {
		x += buttonWidth * 1.1
		y := startY - buttonHeight/2

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
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
