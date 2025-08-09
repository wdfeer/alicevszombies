package internal

import (
	"alicevszombies/internal/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShootPattern struct {
	projectile              *ProjectileType
	velocity                float32
	cooldown                float32
	typ                     ShootType
	count                   uint8
	countExtraPerDifficulty float32
	// Only used on bosses
	countExtraPerWave float32
	// Only used for Spread ShootType
	spread float32
}

type ShootType = uint8

const (
	DirectShoot ShootType = iota
	CircleShoot
	SpreadShoot
	UnknownShoot
)

func updateShooting(world *World) {
	for id := range world.shootTimer {
		world.shootTimer[id] -= dt
		if world.shootTimer[id] < 0 {
			var pattern ShootPattern
			var friendly bool

			if typ, ok := world.enemy[id]; ok {
				pattern = typ.shootPattern
				friendly = false
			} else if typ, ok := world.doll[id]; ok {
				pattern = typ.shootPattern
				friendly = true
			} else {
				println("WARNING: entity with shootTimer is neither an enemy nor a doll!")
				continue
			}

			if e, ok := world.enemy[id]; ok && e.spawnData.boss && (world.hp[id].val/world.hp[id].max) < 0.4 {
				world.shootTimer[id] = pattern.cooldown / 1.5
			} else {
				world.shootTimer[id] = pattern.cooldown
			}

			var target rl.Vector2
			if friendly {
				enemyFound := false
				var enemyTarget Entity
				var minDist float32 = 220

				// Prioritize special enemies close to player
				for enemy, typ := range world.enemy {
					if typ != &enemyTypes.zombie {
						ownDist := rl.Vector2Distance(world.position[id], world.position[enemy])
						playerDist := rl.Vector2Distance(world.position[world.player], world.position[enemy])
						if ownDist < minDist && playerDist < minDist/2 {
							enemyFound = true
							enemyTarget = enemy
							minDist = ownDist
						}
					}
				}

				if !enemyFound {
					for enemy := range world.enemy {
						dist := rl.Vector2Distance(world.position[id], world.position[enemy])
						if dist < minDist {
							enemyFound = true
							enemyTarget = enemy
							minDist = dist
						}
					}
				}

				if enemyFound {
					target = world.position[enemyTarget]
				} else {
					continue
				}
			} else {
				target = world.position[world.player]
			}

			dir := util.Vector2Direction(world.position[id], target)
			vel := rl.Vector2Scale(dir, pattern.velocity)

			switch pattern.typ {
			case DirectShoot:
				newProjectile(world, world.position[id], vel, pattern.projectile)
			case CircleShoot:
				count := pattern.count + uint8(pattern.countExtraPerWave*float32(world.enemySpawner.wave)+pattern.countExtraPerDifficulty*float32(world.difficulty))
				for i := range count {
					ratio := (float32(i) + 1) / float32(count)
					newProjectile(world, world.position[id], rl.Vector2Rotate(vel, math.Pi*2*ratio), pattern.projectile)
				}
			case SpreadShoot:
				count := pattern.count + uint8(pattern.countExtraPerWave*float32(world.enemySpawner.wave)+pattern.countExtraPerDifficulty*float32(world.difficulty))
				vel := rl.Vector2Rotate(vel, -pattern.spread/2)
				for i := range count {
					ratio := (float32(i) + 1) / float32(count)
					newProjectile(world, world.position[id], rl.Vector2Rotate(vel, pattern.spread*ratio), pattern.projectile)
				}
			case UnknownShoot:
				count := pattern.count + uint8(pattern.countExtraPerWave*float32(world.enemySpawner.wave)+pattern.countExtraPerDifficulty*float32(world.difficulty))
				count += uint8(rand.Float32() * float32(count))
				for range count {
					vel = rl.Vector2Scale(rl.Vector2Add(vel, rl.Vector2Scale(util.Vector2Random(), pattern.velocity/10)), 0.75+rand.Float32()/2)

					// Leading shot on X
					if rand.Float32() < 0.1 {
						if world.velocity[world.player].X < 0 {
							vel.X -= pattern.velocity / 4
						} else {
							vel.X += pattern.velocity / 4
						}
					}

					// Leading shot on Y
					if rand.Float32() < 0.1 {
						if world.velocity[world.player].Y < 0 {
							vel.Y -= pattern.velocity / 4
						} else {
							vel.Y += pattern.velocity / 4
						}
					}

					var projType *ProjectileType
					if rand.Float32() < 0.5 {
						projType = &projectileTypes.purpleBullet
					} else {
						projType = &projectileTypes.blueBullet
					}

					newProjectile(world, world.position[id], vel, projType)
				}
			}
		}
	}
}
