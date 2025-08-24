package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DeathEffectType = uint8

const (
	deathEffectNormal DeathEffectType = iota
	deathEffectExplode
	deathEffectSlice
)

func newDeathEffect(world *World, name string, center rl.Vector2, typ DeathEffectType) {
	particles := newBreakdown(world, name, center)

	sliceVel := rl.Vector2Rotate(rl.Vector2{X: 25, Y: 0}, rand.Float32()*math.Pi)
	minusSliceVel := rl.Vector2Scale(sliceVel, -1)

	for _, id := range particles {
		switch typ {
		case deathEffectNormal:
			world.velocity[id] = rl.Vector2Rotate(rl.Vector2{X: 0, Y: 20}, rand.Float32()-0.5)
		case deathEffectExplode:
			world.velocity[id] = rl.Vector2Scale(util.Vector2Random(), 30)
		case deathEffectSlice:
			dot := rl.Vector2DotProduct(rl.Vector2Subtract(world.position[id], center), sliceVel)
			if math.Signbit(float64(dot)) {
				world.velocity[id] = sliceVel
			} else {
				world.velocity[id] = minusSliceVel
			}
		}

		world.drag[id] = rand.Float32()/10 + 0.1
	}
}
