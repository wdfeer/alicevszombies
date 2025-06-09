package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateVelocity(world *World) {
	for entity, pos := range world.positions {
		if vel, ok := world.velocities[entity]; ok {
			world.positions[entity] = rl.Vector2Add(pos, vel)
		}
	}
}
