package internal

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateDrag(world *World) {
	for id, drag := range world.drag {
		if world.status[id][Slow] > 0 {
			drag += 0.5 + float32(world.difficulty)/3
		}

		if vel, ok := world.velocity[id]; ok {
			speed := rl.Vector2Length(vel)
			if speed > 1 {
				decay := float32(math.Exp(-float64(drag * dt)))
				world.velocity[id] = rl.Vector2Scale(vel, decay)
			} else {
				world.velocity[id] = rl.Vector2Zero()
			}
		}
	}
}
