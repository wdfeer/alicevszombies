package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func UpdateAnimationData(world *World) {
	updatePlayerTexture(world)
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
			world.animTimer[world.player] = world.animTimer[world.player] + rl.GetFrameTime()
		}
	} else {
		world.animTimer[world.player] = 0
		world.texture[world.player] = "player"
	}
}
