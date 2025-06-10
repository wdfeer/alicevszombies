package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type dollData struct {
	target         rl.Vector2
	targetingTimer float32
}

func newDoll(world *World) Entity {
	id := world.newEntity()
	world.doll[id] = dollData{}
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 1
	world.texture[id] = "doll1"
	return id
}

func updateDolls(world *World) {
	for id := range world.doll {
		world.doll[id] = updateTargetingData(world, id)

		delta := rl.Vector2Subtract(world.doll[id].target, world.position[id])
		dir := rl.Vector2Normalize(delta)
		world.velocity[id] = rl.Vector2Add(world.velocity[id], rl.Vector2Scale(dir, 350*rl.GetFrameTime()))
	}
}

func updateTargetingData(world *World, id Entity) dollData {
	doll := world.doll[id]
	doll.targetingTimer += rl.GetFrameTime()
	if doll.targetingTimer > 0.4 || rl.Vector2Distance(doll.target, world.position[id]) < 2 {
		doll.targetingTimer = 0

		plPos := world.position[world.player]
		delta := rl.Vector2Rotate(rl.Vector2{X: 20, Y: 0}, rand.Float32()*math.Pi*2)
		doll.target = rl.Vector2Add(plPos, delta)
	}
	return doll
}

func renderDolls(world *World) {
	for id := range world.doll {
		util.DrawTextureCentered(assets.Textures[world.texture[id]], world.position[id])
	}
}
