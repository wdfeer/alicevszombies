package internal

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateDrag(world *World) {
	for entity, drag := range world.drag {
		if vel, ok := world.velocity[entity]; ok {
			speed := rl.Vector2Length(vel)
			if speed > 1 {
				decay := float32(math.Exp(-float64(drag * dt)))
				world.velocity[entity] = rl.Vector2Scale(vel, decay)
			} else {
				world.velocity[entity] = rl.Vector2Zero()
			}
		}
	}
}
