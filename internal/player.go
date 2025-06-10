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

func renderPlayer(world *World) {
	util.DrawTextureCentered(assets.Textures[world.texture[world.player]], world.position[world.player])
}
