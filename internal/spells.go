package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderSpells(world *World) {
	halfHeight := float32(rl.GetScreenHeight() / 2)
	size := rl.Vector2{X: 200, Y: 80}
	center := rl.Vector2{X: 300, Y: halfHeight - size.Y*1.2}

	if raygui.Button(util.CenterRectangle(center, size), "") && world.playerData.mana >= 5 {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	util.DrawTextureCenteredScaled(assets.textures["heal_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("H", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})

	center.Y += size.Y * 1.2
	if raygui.Button(util.CenterRectangle(center, size), "") && world.playerData.mana >= 10 {
		id := newDoll(world, &dollTypes.basicDoll)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	util.DrawTextureCenteredScaled(assets.textures["doll_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("J", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})

	center.Y += size.Y * 1.2
	if raygui.Button(util.CenterRectangle(center, size), "") && world.playerData.mana >= 10 {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
	util.DrawTextureCenteredScaled(assets.textures["pitem_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("K", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})
}
