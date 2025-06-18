package internal

import (
	"alicevszombies/internal/util"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Both update and render
func renderMainMenu(world *World) {
	screenSize := util.ScreenSize()
	rl.DrawRectangleV(rl.Vector2Zero(), screenSize, rl.ColorAlpha(rl.Black, 0.9))

	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{
		X: screenSize.X / 2,
		Y: 120,
	}, 16)

	buttonWidth := float32(240)
	buttonHeight := float32(60)
	buttonSpacing := float32(20)
	startY := screenSize.Y/2 - (buttonHeight*2 + buttonSpacing*1.5)

	x := screenSize.X/2 - buttonWidth/2
	y := startY

	if rg.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Easy") {
		startGame(world, EASY)
	}
	y += buttonHeight + buttonSpacing

	if rg.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Normal") {
		startGame(world, NORMAL)
	}
	y += buttonHeight + buttonSpacing

	if rg.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Hard") {
		startGame(world, HARD)
	}
	y += buttonHeight + buttonSpacing

	if rg.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Lunatic") {
		startGame(world, LUNATIC)
	}

	bottomY := screenSize.Y - 80
	leftX := float32(80)
	rightX := screenSize.X - buttonWidth - 80

	if rg.Button(rl.Rectangle{X: leftX, Y: bottomY, Width: buttonWidth, Height: buttonHeight}, "Toggle Fullscreen") {
		rl.ToggleFullscreen()
	}

	if rg.Button(rl.Rectangle{X: rightX, Y: bottomY, Width: buttonWidth, Height: buttonHeight}, "Exit Game") {
		rl.CloseWindow()
	}
}

func startGame(world *World, difficulty Difficulty) {
	world.difficulty = difficulty
	world.paused = false
	world.uistate.isMainMenu = false
}
