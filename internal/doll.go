package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func spawnDollWithAnimation(world *World, typ *DollType) {
	world.playerData.dollSpawnTimer = 0.5
	world.playerData.dollToSpawn = typ
	world.playerData.dollSpawnPosition = rl.Vector2Add(world.position[world.player], rl.Vector2Scale(util.Vector2Random(), 24))

	particles := newBreakdown(world, typ.texture, world.playerData.dollSpawnPosition)
	for _, id := range particles {
		offset := util.Vector2Random()
		offset = rl.Vector2Scale(offset, 3)

		world.position[id] = rl.Vector2Add(world.position[id], offset)
		world.velocity[id] = rl.Vector2Scale(offset, -1/world.playerData.dollSpawnTimer)
		world.pixelParticle[id] = PixelParticle{
			timeleft:     world.playerData.dollSpawnTimer,
			tint:         world.pixelParticle[id].tint,
			reverseAlpha: true,
		}
	}
}

func newDoll(world *World, typ *DollType) Entity {
	id := world.newEntity()
	world.doll[id] = typ
	world.targeting[id] = Targeting{
		accel: typ.accel,
	}
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 5
	world.size[id] = typ.size
	world.texture[id] = typ.texture
	if typ.contactDamage > 0 {
		world.flippable[id] = true
	}

	if !world.uistate.isMainMenu && !world.uistate.isDeathScreen {
		stats.DollsSummoned[world.difficulty]++
	}

	if typ.shootPattern.projectile != nil {
		world.shootTimer[id] = typ.shootPattern.cooldown
	}

	return id
}

func updateDolls(world *World) {
	for doll := range world.doll {
		world.targeting[doll] = updateDollTargeting(world, doll)
	}
}

func updateDollTargeting(world *World, doll Entity) Targeting {
	typ := world.doll[doll]
	targeting := world.targeting[doll]
	targeting.accel = typ.accel + float32(30*math.Sqrt(float64(world.playerData.upgrades[&DollSpeed])))
	targeting.targetingTimer -= dt
	if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[doll]) < 2 {
		targeting.targetingTimer = 0.4

		var target rl.Vector2

		nextIndex := 0
		validEnemies := [16]Entity{}
		for enemy := range world.enemy {
			dist := rl.Vector2Distance(world.position[doll], world.position[enemy])
			playerDist := rl.Vector2Distance(world.position[world.player], world.position[enemy])

			if playerDist < 70 && doll%3 < 2 {
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
			enemyPos := world.position[enemy]
			if typ.contactDamage > 0 {
				target = enemyPos
			} else {
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
