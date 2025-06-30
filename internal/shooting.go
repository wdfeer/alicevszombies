package internal

import (
	"alicevszombies/internal/util"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShootPattern struct {
	projectile *ProjectileType
	velocity   float32
	cooldown   float32
	typ        ShootType
	count      uint8
	// Only used on bosses
	countExtraPerWave float32
	// Only used for Spread ShootType
	spread float32
}

type ShootType = uint8

const (
	Direct ShootType = iota
	Circle
	Spread
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

			world.shootTimer[id] = pattern.cooldown

			var target rl.Vector2
			if friendly {
				enemyFound := false
				var enemyTarget Entity
				var minDist float32 = 220
				for enemy := range world.enemy {
					dist := rl.Vector2Distance(world.position[id], world.position[enemy])
					if dist < minDist {
						enemyFound = true
						enemyTarget = enemy
						minDist = dist
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
			case Direct:
				newProjectile(world, world.position[id], vel, pattern.projectile)
			case Circle:
				count := pattern.count + uint8(pattern.countExtraPerWave*float32(world.enemySpawner.wave))
				for i := range count {
					ratio := (float32(i) + 1) / float32(count)
					newProjectile(world, world.position[id], rl.Vector2Rotate(vel, math.Pi*2*ratio), pattern.projectile)
				}
			case Spread:
				count := pattern.count + uint8(pattern.countExtraPerWave*float32(world.enemySpawner.wave))
				vel := rl.Vector2Rotate(vel, -pattern.spread/2)
				for i := range count {
					ratio := (float32(i) + 1) / float32(count)
					newProjectile(world, world.position[id], rl.Vector2Rotate(vel, pattern.spread*ratio), pattern.projectile)
				}
			}
		}
	}
}
