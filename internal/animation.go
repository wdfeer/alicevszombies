package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateAnimationData(world *World) {
	updateWalkAnimations(world)
	updateDollTexture(world)
}

type WalkAnimation struct {
	baseTexture string
}

func updateWalkAnimations(world *World) {
	for id, data := range world.walkAnimated {
		if _, ok := world.velocity[id]; !ok {
			return
		}

		if rl.Vector2Length(world.velocity[id]) > 0 {
			if world.animTimer[id] > 0.15 {
				world.animTimer[id] = 0
				if world.texture[id] == data.baseTexture+"_walk0" {
					world.texture[id] = data.baseTexture + "_walk1"
				} else {
					world.texture[id] = data.baseTexture + "_walk0"
				}
			} else {
				world.animTimer[id] = world.animTimer[id] + dt
			}
		} else {
			world.animTimer[id] = 0
			world.texture[id] = data.baseTexture
		}
	}
}

func updateDollTexture(world *World) {
	for id := range world.dollTag {
		if world.velocity[id].X >= 0 {
			world.texture[id] = "doll"
		} else {
			world.texture[id] = "doll_fliph"
		}
	}
}
