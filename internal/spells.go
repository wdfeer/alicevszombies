package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateSpells(world *World) {
	if world.playerData.mana >= 5 && (rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyH)) {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyJ)) {
		id := newDoll(world, &dollTypes.basicDoll)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyThree) || rl.IsKeyPressed(rl.KeyK)) {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
}

func renderSpellsBar() {
	util.DrawTextureCenteredScaled(assets.textures["heal_icon"],
		rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 - 80},
		4)
	util.DrawTextCentered("H", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight())/2 - 80})
	util.DrawTextCentered("5 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 - 50})
	util.DrawTextureCenteredScaled(assets.textures["doll_icon"],
		rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight()) / 2},
		4)
	util.DrawTextCentered("J", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight()) / 2})
	util.DrawTextCentered("10 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 30})
	util.DrawTextureCenteredScaled(assets.textures["pitem_icon"],
		rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 80},
		4)
	util.DrawTextCentered("K", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight())/2 + 80})
	util.DrawTextCentered("10 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 110})
}
