package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateSpells(world *World) {
	if world.playerData.mana >= 10 {
		if rl.IsKeyPressed(rl.KeyOne) {
			id := newDoll(world)
			world.position[id] = world.position[world.player]
			world.playerData.mana -= 10
		} else if rl.IsKeyPressed(rl.KeyTwo) {
			heal(world, world.player, 5)
			world.playerData.mana -= 10
		}
	}
}
