package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderMainMenu(world *World) {
	screenSize := util.ScreenSize()
	rl.DrawRectangleV(rl.Vector2Zero(), screenSize, rl.ColorAlpha(rl.Black, 0.9))

	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{
		X: screenSize.X / 2,
		Y: 120,
	}, 16)

	buttonWidth := float32(400)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)
	startY := screenSize.Y/2 - (buttonHeight*2 + buttonSpacing*1.5)

	x := screenSize.X/2 - buttonWidth/2
	y := startY

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 80)
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

	bottomY := screenSize.Y - buttonHeight - 80
	leftX := float32(80)
	rightX := screenSize.X - buttonWidth - 80

	if raygui.Button(rl.Rectangle{X: leftX, Y: bottomY, Width: buttonWidth, Height: buttonHeight}, "Fullscreen") {
		rl.ToggleFullscreen()
	}

	if raygui.Button(rl.Rectangle{X: rightX, Y: bottomY, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
