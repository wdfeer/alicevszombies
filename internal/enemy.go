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
		if world.enemySpawner.wave > 60 {
			hp *= float32(math.Exp(float64(world.enemySpawner.wave) / 60))
		}
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
	Confused
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
				delta := util.Vector2Direction(world.position[id], world.position[world.player])
				delta = rl.Vector2Rotate(delta, rand.Float32()/2)
				delta = rl.Vector2Scale(delta, distance/3)
				targeting.target = rl.Vector2Add(world.position[id], delta)
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
			case Confused:
				delta := util.Vector2Direction(world.position[id], world.position[world.player])
				delta = rl.Vector2Rotate(delta, rand.Float32()/2)
				delta = rl.Vector2Scale(delta, distance/3)
				delta = rl.Vector2Rotate(delta, (rand.Float32()-0.5)*2)
				targeting.target = rl.Vector2Add(world.position[id], delta)
			case Ranged:
				targeting.target = rl.Vector2Add(world.position[world.player], rl.Vector2Scale(util.Vector2Random(), 70))
			}
		}
		world.targeting[id] = targeting

		if typ.spawnData.boss {
			const immuneDuration = 10
			newHP := world.hp[id]
			ratio := newHP.val / newHP.max
			if ratio > 0.4 && ratio < 0.6 {
				dps := 0.2 / immuneDuration * newHP.max
				newHP.val -= dps * dt
				world.hp[id] = newHP
			}
		}
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

	deathEffectType := deathEffectNormal
	var attacker Entity
	maxCooldown := float32(0)
	for id, v := range world.hp[id].attackerCooldown {
		if v > maxCooldown {
			attacker = id
			maxCooldown = v
		}
	}
	if dollType, ok := world.doll[attacker]; ok {
		switch dollType {
		case &dollTypes.scytheDoll:
			deathEffectType = deathEffectSlice
		case &dollTypes.destructionDoll:
			deathEffectType = deathEffectExplode
		}
	} else if proj, ok := world.projectile[attacker]; ok {
		switch proj.typ {
		case &projectileTypes.knife:
			deathEffectType = deathEffectSlice
		case &projectileTypes.redBullet:
			deathEffectType = deathEffectExplode
		}
	}

	newDeathEffect(world, world.enemy[id].texture, world.position[id], deathEffectType)

	history.EnemiesKilledTotal[world.difficulty]++
	history.EnemiesKilledPerType[world.enemy[id].texture]++

	dist := rl.Vector2Distance(world.position[id], world.position[world.player])
	playSoundVolumePitch("enemy_kill", (0.75 - dist/400), 0.9+0.15*rand.Float32())
}
