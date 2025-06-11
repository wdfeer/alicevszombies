package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateDrag(world *World) {
	for entity, drag := range world.drag {
		if vel, ok := world.velocity[entity]; ok {
			if rl.Vector2Length(world.velocity[entity]) > 1 {
				world.velocity[entity] = rl.Vector2Scale(vel, 1-drag*dt*400/(drag+100))
			} else {
				world.velocity[entity] = rl.Vector2Zero()
			}
		}
	}
}
