package internal

import (
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
			if rl.Vector2Distance(world.position[doll], world.position[enemy]) < 16 {
				damageWithCooldown(world, enemy, 1+(float32(world.playerData.upgrades[DOLL_DAMAGE])/4), doll)
				break
			}
		}
	}
}

func updateDollTargeting(world *World, doll Entity) Targeting {
	targeting := world.targeting[doll]
	targeting.targetingTimer -= dt
	if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[doll]) < 2 {
		targeting.targetingTimer = 0.4

		var target rl.Vector2
		enemyTargetFound := false

		for enemy := range world.enemyTag {
			dist := rl.Vector2Distance(world.position[doll], world.position[enemy])
			playerDist := rl.Vector2Distance(world.position[world.player], world.position[enemy])
			if dist < 160 && playerDist < 180 {
				target = world.position[enemy]
				enemyTargetFound = true
				break
			}
		}

		if !enemyTargetFound {
			delta := rl.Vector2Rotate(rl.Vector2{X: 20, Y: 0}, rand.Float32()*math.Pi*2)
			target = rl.Vector2Add(world.position[world.player], delta)
		}

		targeting.target = target
	}
	return targeting
}
