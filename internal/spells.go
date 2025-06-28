package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderSpells(world *World) {
	halfHeight := float32(rl.GetScreenHeight() / 2)
	size := rl.Vector2{X: 200, Y: 80}
	pos := rl.Vector2{X: 300, Y: halfHeight}

	if world.playerData.mana < 5 {
		raygui.Disable()
	}

	spellCount := 3
	if superUpgradesAvailable(world) {
		spellCount = 4
	}
	yPositions := util.SpaceCentered(size.Y*1.2, spellCount, pos.Y)

	pos.Y = yPositions[0]
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyH)) && world.playerData.mana >= 5 && !world.paused {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	util.DrawTextureCenteredScaled(assets.textures["heal_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("H", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	if world.playerData.mana < 10 {
		raygui.Disable()
	}

	pos.Y = yPositions[1]
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyJ)) && world.playerData.mana >= 10 && !world.paused {
		id := newDoll(world, &dollTypes.basicDoll)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	util.DrawTextureCenteredScaled(assets.textures["doll_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("J", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	pos.Y = yPositions[2]
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyK)) && world.playerData.mana >= 10 && !world.paused {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
	util.DrawTextureCenteredScaled(assets.textures["pitem_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("K", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	if spellCount == 4 {
		pos.Y = yPositions[3]
		if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyL)) && world.playerData.mana >= 100 && !world.paused {
			world.paused = true
			world.playerData.mana -= 100
			newSuperUpgradeScreen(world)
		}
		// TODO: use a different texture
		util.DrawTextureCenteredScaled(assets.textures["pitem_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
		util.DrawTextCentered("L", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})
	}

	raygui.Enable()
}
