package util

import "math/rand"

// Rounds the float value - the closer integer has a higher probability
func RandomRound(value float32) int {
	if rand.Float32() > value-float32(int(value)) {
		return int(value)
	} else {
		return int(value + 1)
	}
}
