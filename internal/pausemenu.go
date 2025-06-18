package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderPauseMenu(world *World) {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	center := util.HalfScreenSize()
	util.DrawTextCenteredSpaced("Paused", 256, center, 16)

	buttonWidth := float32(400)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)

	startY := center.Y + 128
	centerX := center.X - buttonWidth/2

	resumeRect := rl.Rectangle{X: centerX, Y: startY, Width: buttonWidth, Height: buttonHeight}
	quitRect := rl.Rectangle{X: centerX, Y: startY + buttonHeight + buttonSpacing, Width: buttonWidth, Height: buttonHeight}

	if raygui.Button(resumeRect, "Resume") {
		world.paused = false
	}

	if raygui.Button(quitRect, "Quit") {
		rl.CloseWindow()
	}
}
