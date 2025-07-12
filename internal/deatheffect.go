package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newDeathEffect(world *World, name string, center rl.Vector2) {
	particles := newBreakdown(world, name, center)
	for _, id := range particles {
		world.velocity[id] = rl.Vector2Rotate(rl.Vector2{X: 0, Y: 20}, rand.Float32()-0.5)
		world.drag[id] = rand.Float32()/10 + 0.1
	}
}
