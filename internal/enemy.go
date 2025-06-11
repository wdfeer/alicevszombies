package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var enemySpawnTimer float32 = 2

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
	world.texture[id] = "zombie1"
	world.animTimer[id] = 0
	world.hp[id] = newHP(3)
	return id
}

func updateEnemies(world *World) {
	enemySpawnTimer -= rl.GetFrameTime()
	if enemySpawnTimer <= 0 {
		newEnemy(world)
		enemySpawnTimer = 2
	}

	for id := range world.enemyTag {
		targeting := world.targeting[id]
		targeting.targetingTimer -= rl.GetFrameTime()
		if targeting.targetingTimer <= 0 || rl.Vector2Distance(targeting.target, world.position[id]) < 2 {
			targeting.targetingTimer = 0.4
			targeting.target = world.position[world.player]
		}
		world.targeting[id] = targeting
	}
}
