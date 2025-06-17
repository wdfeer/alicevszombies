package internal

import (
	"alicevszombies/internal/util"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemySpawner struct {
	wave           uint32
	enemiesToSpawn uint32
	spawnTimer     float32
}

func newEnemySpawner() EnemySpawner {
	return EnemySpawner{
		wave:           0,
		enemiesToSpawn: 0,
		spawnTimer:     0,
	}
}

func newEnemy(world *World, typ *EnemyType) Entity {
	id := world.newEntity()
	world.enemy[id] = typ
	world.targeting[id] = Targeting{
		accel: typ.acceleration,
	}
	world.position[id] = rl.Vector2Add(
		world.position[world.player],
		rl.Vector2Scale(util.Vector2Random(), 500),
	)
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 10
	world.walkAnimated[id] = WalkAnimation{typ.texture}
	world.hp[id] = newHP(typ.baseHP * (1 + float32(world.enemySpawner.wave/20)) * (1 + (float32(world.difficulty)-1)/3))
	world.size[id] = rl.Vector2{X: 8, Y: 16}
	return id
}

func updateEnemySpawner(world *World) {
	spawner := world.enemySpawner

	if spawner.enemiesToSpawn <= 0 {
		spawner.wave++
		spawner.enemiesToSpawn = 1 + spawner.wave*2
	}

	spawner.spawnTimer = spawner.spawnTimer - dt
	if spawner.spawnTimer <= 0 {
		if spawner.wave%10 == 0 && spawner.enemiesToSpawn > 1 {
			newEnemy(world, &enemyTypes.medicine)
			spawner.spawnTimer = 15
			spawner.enemiesToSpawn = 1
		} else {
			newEnemy(world, &enemyTypes.zombie)
			spawner.spawnTimer = 2 - min(1.4, float32(spawner.wave)/10)
			if spawner.enemiesToSpawn > 10 {
				spawner.spawnTimer /= 2
			}
			if spawner.enemiesToSpawn > 30 {
				spawner.spawnTimer /= 2
			}
			spawner.enemiesToSpawn--
		}
	}

	world.enemySpawner = spawner
}

func updateEnemies(world *World) {
	for id, typ := range world.enemy {
		targeting := world.targeting[id]
		targeting.targetingTimer -= dt
		if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[id]) < 2 {
			targeting.targetingTimer = 0.4
			distance := rl.Vector2Distance(world.position[id], world.position[world.player])
			if !typ.ranged {
				delta := rl.Vector2Normalize(rl.Vector2Subtract(world.position[world.player], world.position[id]))
				delta = rl.Vector2Rotate(delta, rand.Float32()/2)
				delta = rl.Vector2Scale(delta, distance/3)
				targeting.target = rl.Vector2Add(world.position[id], delta)
			} else {
				targeting.target = rl.Vector2Add(world.position[world.player], rl.Vector2Scale(util.Vector2Random(), 70))
			}
		}
		world.targeting[id] = targeting

		if typ.ranged {
			world.shootTimer[id] -= dt
			if world.shootTimer[id] <= 0 {
				world.shootTimer[id] = 1

				vel := rl.Vector2Scale(util.Vector2Direction(world.position[id], world.position[world.player]), 100)
				newProjectile(world, world.position[id], vel, &projectileTypes.redBullet)
			}
		}
	}
}

func preEnemyDeath(world *World, entity Entity) {
	world.playerData.mana += 1
	newDeathEffect(world, world.enemy[entity].texture, world.position[entity])
}
