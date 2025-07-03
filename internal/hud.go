package internal

import (
	"alicevszombies/internal/colors"
	"alicevszombies/internal/util"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderHUD(world *World) {
	size := util.ScreenSize()
	halfSize := util.HalfScreenSize()

	// Upper Center

	{ // Wave counter
		str := "Wave " + fmt.Sprint(world.enemySpawner.wave)
		center := rl.Vector2{X: halfSize.X, Y: 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	// Lower Center
	yPositions := util.SpaceCentered(50, 5, float32(rl.GetScreenHeight())-200)

	// Stamina bar
	if world.playerData.stamina < 1 {
		width := float32(200) * world.playerData.stamina
		pos := rl.Vector2{X: halfSize.X - width/2, Y: yPositions[0]}
		size := rl.Vector2{X: width, Y: 8}
		rl.DrawRectangleV(pos, size, colors.Yellow)
	}

	{ // HP
		str := "HP: " + fmt.Sprint(world.hp[world.player].val)
		center := rl.Vector2{X: halfSize.X, Y: yPositions[1]}
		pos := util.CenterText(str, 32, center)
		if world.hp[world.player].val <= 5 {
			pos := util.CenterText(str, 34, center)
			rl.DrawTextEx(rl.GetFontDefault(), str, pos, 34, 2, rl.Red)
		}
		rl.DrawTextEx(rl.GetFontDefault(), str, pos, 32, 2, rl.White)
	}

	{ // MP
		str := "MP: " + fmt.Sprint(world.playerData.mana)
		center := rl.Vector2{X: halfSize.X, Y: yPositions[2]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	// Statuses
	if world.status[world.player].poison > 0 {
		str := "Poisoned"
		center := rl.Vector2{X: halfSize.X, Y: yPositions[3]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, colors.Purple)
	}
	if world.status[world.player].slow > 0 {
		str := "Slowed"
		center := rl.Vector2{X: halfSize.X, Y: yPositions[4]}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, colors.Blue)
	}

	{ // Boss Bar
		boss := -1
		for id, typ := range world.enemy {
			if typ.spawnData.boss {
				boss = int(id)
				break
			}
		}
		if boss != -1 {
			hp := world.hp[Entity(boss)]
			width := int32(size.X * hp.val / hp.max)
			rl.DrawRectangle(0, int32(size.Y-16), width, 16, colors.Red)
		}
	}

	renderSpells(world)
}
