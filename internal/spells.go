package internal

import (
	"alicevszombies/internal/ui"
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderSpells(world *World) {
	halfHeight := float32(rl.GetScreenHeight() / 2)
	size := rl.Vector2{X: 200, Y: 80}
	center := rl.Vector2{X: 300, Y: halfHeight - size.Y*1.2}

	if world.playerData.mana < 5 {
		raygui.Disable()
	}

	if (ui.Button(util.CenterRectangle(center, size), "") || rl.IsKeyPressed(rl.KeyH)) && world.playerData.mana >= 5 && !world.paused {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	ui.TextureC(assets.textures["heal_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("H", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})

	if world.playerData.mana < 10 {
		raygui.Disable()
	}

	center.Y += size.Y * 1.2
	if (ui.Button(util.CenterRectangle(center, size), "") || rl.IsKeyPressed(rl.KeyJ)) && world.playerData.mana >= 10 && !world.paused {
		id := newDoll(world, &dollTypes.basicDoll)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	ui.TextureC(assets.textures["doll_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("J", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})

	center.Y += size.Y * 1.2
	if (ui.Button(util.CenterRectangle(center, size), "") || rl.IsKeyPressed(rl.KeyK)) && world.playerData.mana >= 10 && !world.paused {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
	ui.TextureC(assets.textures["pitem_icon"], rl.Vector2{X: center.X - size.X/5, Y: center.Y}, 4)
	util.DrawTextCentered("K", 40, rl.Vector2{X: center.X + size.X/5, Y: center.Y})

	raygui.Enable()
}
