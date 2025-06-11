package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateAnimationData(world *World) {
	updatePlayerTexture(world)
	updateZombieTexture(world)
}

func updatePlayerTexture(world *World) {
	if rl.Vector2Length(world.velocity[world.player]) > 0 {
		if world.animTimer[world.player] > 0.15 {
			world.animTimer[world.player] = 0
			if world.texture[world.player] == "player_walk0" {
				world.texture[world.player] = "player_walk1"
			} else {
				world.texture[world.player] = "player_walk0"
			}
		} else {
			world.animTimer[world.player] = world.animTimer[world.player] + dt
		}
	} else {
		world.animTimer[world.player] = 0
		world.texture[world.player] = "player"
	}
}

func updateZombieTexture(world *World) {
	for id := range world.enemyTag {
		if rl.Vector2Length(world.velocity[id]) > 0 {
			if world.animTimer[id] > 0.15 {
				world.animTimer[id] = 0
				if world.texture[id] == "zombie1_walk0" {
					world.texture[id] = "zombie1_walk1"
				} else {
					world.texture[id] = "zombie1_walk0"
				}
			} else {
				world.animTimer[id] = world.animTimer[id] + dt
			}
		} else {
			world.animTimer[id] = 0
			world.texture[id] = "zombie1"
		}
	}
}
