package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newEnemy(world *World, typ *EnemyType) Entity {
	id := world.newEntity()
	world.enemy[id] = typ

	accel := typ.acceleration
	if world.difficulty == EASY {
		accel *= 0.92
	}
	world.targeting[id] = Targeting{
		accel: accel,
	}

	world.position[id] = rl.Vector2Add(
		world.position[world.player],
		rl.Vector2Scale(util.Vector2Random(), 500),
	)
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = typ.drag

	world.texture[id] = typ.texture
	if !typ.disableWalking {
		world.walkAnimated[id] = WalkAnimation{typ.texture}
	}
	if typ.flippable {
		world.flippable[id] = true
	}
	world.size[id] = typ.size

	hp := typ.baseHP * (1 + float32(world.enemySpawner.wave/(23-uint32(world.difficulty)*3)))
	if world.enemySpawner.wave > 33-uint32(world.difficulty)*3 {
		hp *= 1 + float32(world.enemySpawner.wave-30+uint32(world.difficulty)*5)/30
	}
	world.hp[id] = newHP(hp)

	if typ.shootPattern.projectile != nil {
		world.shootTimer[id] = typ.shootPattern.cooldown
	}

	return id
}

type EnemyTargetingType = uint8

const (
	DirectMelee EnemyTargetingType = iota
	LeadingMelee
	CirclingMelee
	Ranged
)

func updateEnemies(world *World) {
	for id, typ := range world.enemy {
		targeting := world.targeting[id]
		targeting.targetingTimer -= dt
		if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[id]) < 4 {
			targeting.targetingTimer = 0.4

			distance := rl.Vector2Distance(world.position[id], world.position[world.player])
			switch typ.targetingType {
			case DirectMelee:
				dir := util.Vector2Direction(world.position[id], world.position[world.player])
				dir = rl.Vector2Rotate(dir, rand.Float32()/2)
				dir = rl.Vector2Scale(dir, distance/3)
				targeting.target = rl.Vector2Add(world.position[id], dir)
			case LeadingMelee:
				delta := util.Vector2Direction(world.position[id], world.position[world.player])
				delta = rl.Vector2Rotate(delta, rand.Float32()/2)
				delta = rl.Vector2Lerp(delta, rl.Vector2Normalize(world.velocity[world.player]), 0.2)
				delta = rl.Vector2Scale(delta, distance/2)
				targeting.target = rl.Vector2Add(world.position[id], delta)
			case CirclingMelee:
				dir := util.Vector2Direction(world.position[id], world.position[world.player])
				dir = rl.Vector2Rotate(dir, rand.Float32()/2)
				velPart := rl.Vector2Normalize(world.velocity[id])
				velPart = rl.Vector2Scale(velPart, 0.5)
				dir = rl.Vector2Normalize(rl.Vector2Add(dir, velPart))
				targeting.target = rl.Vector2Add(world.position[id], rl.Vector2Scale(dir, distance))
			case Ranged:
				targeting.target = rl.Vector2Add(world.position[world.player], rl.Vector2Scale(util.Vector2Random(), 70))
			}
		}
		world.targeting[id] = targeting
	}
}

func preEnemyDeath(world *World, id Entity) {
	deathExplode := world.enemy[id].deathExplode
	if deathExplode.active {
		count := deathExplode.getProjectileCount(world)
		for i := range count {
			ratio := (float32(i) + 1) / float32(count)
			newProjectile(world, world.position[id], rl.Vector2Rotate(rl.Vector2{X: 80, Y: 0}, math.Pi*2*ratio), deathExplode.projectileType)
		}
	}

	world.playerData.mana += 1
	newDeathEffect(world, world.enemy[id].texture, world.position[id])

	stats.EnemiesKilled[world.difficulty]++
}
