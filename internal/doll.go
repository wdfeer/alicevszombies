package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newDoll(world *World) Entity {
	id := world.newEntity()
	world.dollTag[id] = true
	world.targeting[id] = Targeting{
		accel: 350,
	}
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 1
	world.texture[id] = "doll1"
	return id
}

func updateDolls(world *World) {
	for doll := range world.dollTag {
		world.targeting[doll] = updateDollTargeting(world, doll)

		for enemy := range world.enemyTag {
			if rl.Vector2Distance(world.position[doll], world.position[enemy]) < 24 {
				damageWithCooldown(world, enemy, 1, doll)
			}
		}
	}
}

func updateDollTargeting(world *World, id Entity) Targeting {
	doll := world.targeting[id]
	doll.targetingTimer -= rl.GetFrameTime()
	if doll.targetingTimer <= 0 || rl.Vector2Distance(doll.target, world.position[id]) < 2 {
		doll.targetingTimer = 0.4

		plPos := world.position[world.player]
		delta := rl.Vector2Rotate(rl.Vector2{X: 20, Y: 0}, rand.Float32()*math.Pi*2)
		doll.target = rl.Vector2Add(plPos, delta)
	}
	return doll
}

func renderDolls(world *World) {
	for id := range world.targeting {
		util.DrawTextureCentered(assets.Textures[world.texture[id]], world.position[id])
	}
}
