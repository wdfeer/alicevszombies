package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateSpells(world *World) {
	if world.playerData.mana >= 5 && (rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyH)) {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyJ)) {
		id := newDoll(world)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyThree) || rl.IsKeyPressed(rl.KeyK)) {
		world.paused = true
		uistate.upgradeScreen = true
		world.playerData.mana -= 10
	}
}
