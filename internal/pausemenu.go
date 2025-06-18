package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderPauseMenu(world *World) {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	pos := util.HalfScreenSize()
	util.DrawTextCenteredSpaced("Paused", 256, pos, 16)
	pos.Y += 128
	util.DrawTextCenteredSpaced("ESC = Resume", 64, pos, 4)
	pos.Y += 64
	util.DrawTextCenteredSpaced("DEL = Quit", 64, pos, 4)
}
