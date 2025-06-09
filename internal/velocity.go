package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateVelocity(world *World) {
	for entity, pos := range world.position {
		if vel, ok := world.velocity[entity]; ok {
			delta := rl.Vector2Scale(vel, rl.GetFrameTime())
			world.position[entity] = rl.Vector2Add(pos, delta)
		}
	}
}
