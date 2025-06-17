package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BASE_DOLL_ACCELERATION = 500

type DollType struct {
	contactDamage float32
	texture       string
}

func newDoll(world *World, typ DollType) Entity {
	id := world.newEntity()
	world.doll[id] = typ
	world.targeting[id] = Targeting{
		accel: BASE_DOLL_ACCELERATION,
	}
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 5
	world.size[id] = rl.Vector2{X: 8, Y: 8}
	world.texture[id] = typ.texture
	return id
}

func updateDolls(world *World) {
	for doll, typ := range world.doll {
		world.targeting[doll] = updateDollTargeting(world, doll)
		if typ.contactDamage == 0 {
			updateDollRanged(world, doll)
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

		nextIndex := 0
		validEnemies := [16]Entity{}
		for enemy := range world.enemy {
			dist := rl.Vector2Distance(world.position[doll], world.position[enemy])
			playerDist := rl.Vector2Distance(world.position[world.player], world.position[enemy])

			if playerDist < 60 && doll%2 == 0 {
				targeting.target = world.position[enemy]
				return targeting
			}

			if dist < 160 && playerDist < 180 {
				validEnemies[nextIndex] = enemy
				nextIndex++
			}
			if nextIndex == len(validEnemies)-1 {
				break
			}
		}

		if nextIndex > 0 {
			random := rand.New(rand.NewSource(int64(doll)))
			enemy := validEnemies[random.Int()%nextIndex]
			dollType := world.doll[doll]
			enemyPos := world.position[enemy]
			switch {
			case dollType.contactDamage > 0:
				target = enemyPos
			default:
				target = rl.Vector2Add(enemyPos, rl.Vector2Scale(util.Vector2Random(), 32))
			}
		} else {
			delta := rl.Vector2Rotate(rl.Vector2{X: 20, Y: 0}, rand.Float32()*math.Pi*2)
			target = rl.Vector2Add(world.position[world.player], delta)
		}

		targeting.target = target
	}
	return targeting
}

func updateDollRanged(world *World, doll Entity) {
	world.shootTimer[doll] -= dt
	if world.shootTimer[doll] <= 0 {
		enemyFound := false
		var enemyTarget Entity
		var minDist float32 = 9e25
		for enemy := range world.enemy {
			dist := rl.Vector2Distance(world.position[doll], world.position[enemy])
			if dist < minDist {
				enemyFound = true
				enemyTarget = enemy
				minDist = dist
			}
		}

		if enemyFound {
			dir := util.Vector2Direction(world.position[doll], world.position[enemyTarget])
			newProjectile(world, world.position[doll], rl.Vector2Scale(dir, 200), &projectileTypes.knife)
			world.shootTimer[doll] = 1
		}
	}
}
