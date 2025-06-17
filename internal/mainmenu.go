package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateMainMenu(world *World) {
	if rl.IsKeyPressed(rl.KeyOne) {
		world.paused = false
		world.uistate.isMainMenu = false
	}
}

func renderMainMenu() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.9))
	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 120}, 16)
	// TODO: render title and start text, key press tooltip
}
