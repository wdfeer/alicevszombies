package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderUI(world *World) {
	rl.DrawTextureEx(assets.Textures["cursor"], rl.GetMousePosition(), 0, 4, rl.White)

	{ // Wave counter
		str := "Wave " + fmt.Sprint(world.enemySpawner.wave)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight())/2 - 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	{ // Hp bar
		str := "HP: " + fmt.Sprint(world.hp[world.player].val)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight())/2 + 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}
}
