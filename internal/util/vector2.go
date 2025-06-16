package util

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Returns a normalized vector in a random direction
func Vector2Random() rl.Vector2 {
	return rl.Vector2Rotate(rl.Vector2{X: 1, Y: 0}, math.Pi*rand.Float32()*2)
}
