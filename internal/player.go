package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newPlayer(world *World) Entity {
	world.player = world.newEntity()
	world.position[world.player] = rl.Vector2Zero()
	world.velocity[world.player] = rl.Vector2Zero()
	world.drag[world.player] = 10

	return world.player
}

const PLAYER_ACCELERATION = 3000

func updatePlayer(world *World) {
	dir := rl.Vector2Zero()

	inputMap := map[int32]rl.Vector2{
		rl.KeyW: {X: 0, Y: -1},
		rl.KeyA: {X: -1, Y: 0},
		rl.KeyS: {X: 0, Y: 1},
		rl.KeyD: {X: 1, Y: 0},
	}

	for k, v := range inputMap {
		if rl.IsKeyDown(k) {
			dir = rl.Vector2Add(dir, v)
		}
	}
	dir = rl.Vector2Normalize(dir)

	delta := rl.Vector2Scale(dir, PLAYER_ACCELERATION*rl.GetFrameTime())
	world.velocity[world.player] = rl.Vector2Add(world.velocity[world.player], delta)
}

func renderPlayer(world *World) {
	util.DrawTextureCentered(assets.Textures[world.texture[world.player]], world.position[world.player])
}
