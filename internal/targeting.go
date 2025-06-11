package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Targeting struct {
	target         rl.Vector2
	targetingTimer float32
	accel          float32
}

func updateTargetingMovement(world *World) {
	for id, data := range world.targeting {
		delta := rl.Vector2Subtract(data.target, world.position[id])
		dir := rl.Vector2Normalize(delta)
		world.velocity[id] = rl.Vector2Add(world.velocity[id], rl.Vector2Scale(dir, data.accel*dt))
	}
}
