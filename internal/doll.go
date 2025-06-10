package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newDoll(world *World) Entity {
	id := world.newEntity()
	world.dollTag[id] = true
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	world.texture[id] = "doll1"
	return id
}

func renderDolls(world *World) {
	for id := range world.dollTag {
		util.DrawTextureCentered(assets.Textures[world.texture[id]], world.position[id])
	}
}
