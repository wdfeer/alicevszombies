package internal

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BASE_DOLL_ACCELERATION = 350

func newDoll(world *World) Entity {
	id := world.newEntity()
	world.dollTag[id] = true
	world.targeting[id] = Targeting{
		accel: BASE_DOLL_ACCELERATION,
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
	targeting.accel = float32(BASE_DOLL_ACCELERATION + 10*world.playerData.upgrades[DOLL_SPEED])
	targeting.targetingTimer -= dt
	if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[doll]) < 2 {
		targeting.targetingTimer = 0.4

		var target rl.Vector2
		enemyTargetFound := false

		random := rand.New(rand.NewSource(int64(doll)))
		nextIndex := 0
		validEnemies := [16]Entity{}
		for enemy := range world.enemyTag {
			dist := rl.Vector2Distance(world.position[doll], world.position[enemy])
			playerDist := rl.Vector2Distance(world.position[world.player], world.position[enemy])
			if dist < 160 && playerDist < 180 {
				validEnemies[nextIndex] = enemy
				nextIndex++
				enemyTargetFound = true
			}
			if nextIndex == len(validEnemies)-1 {
				break
			}
		}

		if enemyTargetFound {
			target = world.position[validEnemies[random.Int()%(nextIndex+1)]]
		} else {
			delta := rl.Vector2Rotate(rl.Vector2{X: 20, Y: 0}, rand.Float32()*math.Pi*2)
			target = rl.Vector2Add(world.position[world.player], delta)
		}

		targeting.target = target
	}
	return targeting
}
