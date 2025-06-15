package internal

import (
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

func newEnemy(world *World) Entity {
	id := world.newEntity()
	world.enemyTag[id] = true
	world.targeting[id] = Targeting{
		accel: 700,
	}
	world.position[id] = rl.Vector2Add(
		world.position[world.player],
		rl.Vector2{X: (rand.Float32() - 0.5) * 500, Y: (rand.Float32() - 0.5) * 500},
	)
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 10
	world.walkAnimated[id] = WalkAnimation{"zombie"}
	world.hp[id] = newHP(3 + float32(world.enemySpawner.wave/10))
	world.size[id] = rl.Vector2{X: 8, Y: 16}
	return id
}

func newMedicine(world *World) Entity {
	id := world.newEntity()
	world.enemyTag[id] = true
	world.targeting[id] = Targeting{
		accel: 730,
	}
	world.position[id] = rl.Vector2Add(
		world.position[world.player],
		rl.Vector2{X: (rand.Float32() - 0.5) * 1000, Y: (rand.Float32() - 0.5) * 1000},
	)
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 10
	world.walkAnimated[id] = WalkAnimation{"medicine"}
	world.hp[id] = newHP(60 + 4*float32(world.enemySpawner.wave))
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
			newMedicine(world)
			spawner.spawnTimer = 15
			spawner.enemiesToSpawn = 1
		} else {
			newEnemy(world)
			spawner.spawnTimer = 2 - min(1, float32(spawner.wave)/10)
			spawner.enemiesToSpawn--
		}
	}

	world.enemySpawner = spawner
}

func updateEnemies(world *World) {
	for id := range world.enemyTag {
		targeting := world.targeting[id]
		targeting.targetingTimer -= dt
		if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[id]) < 2 {
			targeting.targetingTimer = 0.4
			distance := rl.Vector2Distance(world.position[id], world.position[world.player])
			if distance > 10 {
				delta := rl.Vector2Normalize(rl.Vector2Subtract(world.position[world.player], world.position[id]))
				delta = rl.Vector2Rotate(delta, rand.Float32()/2)
				delta = rl.Vector2Scale(delta, distance/3)
				targeting.target = rl.Vector2Add(world.position[id], delta)
			} else {
				targeting.target = world.position[world.player]
			}
		}
		world.targeting[id] = targeting
	}
}
