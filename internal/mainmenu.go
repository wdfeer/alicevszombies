package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	difficultySelect bool
	options          bool
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
			world.uistate.mainMenu.options = false
		}
	}

	y += buttonHeight + buttonSpacing

	if world.uistate.mainMenu.options {
		raygui.SetState(raygui.STATE_PRESSED)
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Options") {
			world.uistate.mainMenu.options = !world.uistate.mainMenu.options
		}
		raygui.SetState(raygui.STATE_NORMAL)
	} else {
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Options") {
			world.uistate.mainMenu.options = !world.uistate.mainMenu.options
			world.uistate.mainMenu.difficultySelect = false
		}
	}

	y += buttonHeight + buttonSpacing

	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}

	x += buttonWidth * 1.1
	y = startY - buttonHeight/2
	if world.uistate.mainMenu.difficultySelect {
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
	} else if world.uistate.mainMenu.options {
		var maxTextWidth float32
		volumeTextSize := float32(rl.MeasureText("Volume", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
		cursorTextSize := float32(rl.MeasureText("Cursor", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
		maxTextWidth = max(volumeTextSize, cursorTextSize)

		buttonWidth -= maxTextWidth / 2
		soundVolume = raygui.SliderBar(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "", "Volume", soundVolume, 0, 1)

		y += buttonHeight + buttonSpacing
		raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(buttonWidth)/7)
		raygui.Spinner(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Cursor", &world.uistate.cursorType, 0, 1, false)

		y += buttonHeight + buttonSpacing
		buttonWidth += maxTextWidth
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen") {
			rl.ToggleFullscreen()
		}
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
