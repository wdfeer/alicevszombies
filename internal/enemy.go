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
		accel: 1500,
	}
	world.position[id] = rl.Vector2Add(
		world.position[world.player],
		rl.Vector2{X: (rand.Float32() - 0.5) * 500, Y: (rand.Float32() - 0.5) * 500},
	)
	world.velocity[id] = rl.Vector2Zero()
	world.drag[id] = 6
	world.texture[id] = "zombie"
	world.animTimer[id] = 0
	world.hp[id] = newHP(3)
	world.size[id] = rl.Vector2{}
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
		newEnemy(world)
		spawner.spawnTimer = 2 - min(1, float32(spawner.wave)/10)
		spawner.enemiesToSpawn--
	}

	world.enemySpawner = spawner
}

func updateEnemies(world *World) {
	for id := range world.enemyTag {
		targeting := world.targeting[id]
		targeting.targetingTimer -= dt
		if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[id]) < 2 {
			targeting.targetingTimer = 0.4
			targeting.target = world.position[world.player]
		}
		world.targeting[id] = targeting
	}
}
