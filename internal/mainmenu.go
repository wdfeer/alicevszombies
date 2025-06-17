package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateMainMenu(world *World) {
	if rl.IsKeyPressed(rl.KeyJ) {
		world.paused = false
		world.uistate.isMainMenu = false
	}
}

func renderMainMenu() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.9))
	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 120}, 16)

	halfHeight := int32(rl.GetScreenHeight() / 2)
	rl.DrawText("Start", 160, halfHeight-50, 64, rl.White)
	rl.DrawText("J", 400, halfHeight-50, 60, rl.White)
	rl.DrawText("Exit", 160, halfHeight+50, 64, rl.White)
	rl.DrawText("DEL", 400, halfHeight+50, 60, rl.White)
}
