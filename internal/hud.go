package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderHUD(world *World) {
	// Top
	{ // Wave counter
		str := "Wave " + fmt.Sprint(world.enemySpawner.wave)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	// Bottom
	yPositions := util.SpaceCentered(50, 4, float32(rl.GetScreenHeight())-200)
	{ // HP bar
		str := "HP: " + fmt.Sprint(world.hp[world.player].val)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: yPositions[0]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	{ // MP bar
		str := "MP: " + fmt.Sprint(world.playerData.mana)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: yPositions[1]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	// Statuses
	if world.status[world.player].poison > 0 {
		str := "Poisoned"
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: yPositions[2]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.Purple)
	}
	if world.status[world.player].slow > 0 {
		str := "Slow"
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: yPositions[3]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.Blue)
	}

	renderSpells(world)
}
