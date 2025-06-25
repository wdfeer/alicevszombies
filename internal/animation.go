package internal

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateAnimationData(world *World) {
	updateWalkAnimations(world)
	updateFlipping(world)
}

type WalkAnimation struct {
	baseTexture string
}

func updateWalkAnimations(world *World) {
	for id, data := range world.walkAnimated {
		if _, ok := world.velocity[id]; !ok {
			return
		}

		var newTexture string
		if rl.Vector2Length(world.velocity[id]) > 0 {
			if world.animTimer[id] > 0.15 {
				world.animTimer[id] = 0
				if strings.HasSuffix(world.texture[id], "_walk0") || strings.HasSuffix(world.texture[id], "_walk0"+FlippedSuffix) {
					newTexture = data.baseTexture + "_walk1"
				} else {
					newTexture = data.baseTexture + "_walk0"
				}
			} else {
				world.animTimer[id] += dt
				continue
			}
		} else {
			world.animTimer[id] = 0
			newTexture = data.baseTexture
		}

		if strings.HasSuffix(world.texture[id], FlippedSuffix) {
			newTexture += FlippedSuffix
		}

		world.texture[id] = newTexture
	}
}

func updateFlipping(world *World) {
	for id := range world.flippable {
		texture := world.texture[id]
		isFlipped := strings.HasSuffix(texture, FlippedSuffix)

		if world.velocity[id].X >= 0 {
			if isFlipped {
				world.texture[id] = strings.TrimSuffix(texture, FlippedSuffix)
			}
		} else {
			if !isFlipped {
				world.texture[id] = texture + FlippedSuffix
			}
		}
	}
}
