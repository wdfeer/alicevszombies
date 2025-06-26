package internal

import (
	"alicevszombies/internal/util"
	"math"
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
	world.drag[id] = 10
	world.walkAnimated[id] = WalkAnimation{typ.texture}
	if typ.flippable {
		world.flippable[id] = true
	}
	world.size[id] = typ.size

	hp := typ.baseHP * (1 + float32(world.enemySpawner.wave/(23-uint32(world.difficulty)*3)))
	if world.enemySpawner.wave > 30 {
		hp *= 1 + float32(world.enemySpawner.wave-30+uint32(world.difficulty)*4)/30
	}
	world.hp[id] = newHP(hp)

	return id
}

func updateEnemySpawner(world *World) {
	spawner := world.enemySpawner

	if spawner.enemiesToSpawn <= 0 {
		spawner.wave++
		spawner.enemiesToSpawn = 2 + spawner.wave*2
	}

	spawner.spawnTimer = spawner.spawnTimer - dt
	if spawner.spawnTimer <= 0 {
		if spawner.wave%10 == 0 && spawner.enemiesToSpawn > 1 {
			switch rand.Int() % 2 {
			case 0:
				newEnemy(world, &enemyTypes.medicine)
			case 1:
				newEnemy(world, &enemyTypes.kogasa)
			}
			spawner.spawnTimer = 15 - float32(world.difficulty)*3
			spawner.enemiesToSpawn = 1
		} else {
			newEnemy(world, enemyTypeToSpawn(world))

			spawner.spawnTimer = 2 - min(1.4, float32(spawner.wave)/10)
			if spawner.enemiesToSpawn > 10 {
				spawner.spawnTimer /= max(2, float32(world.difficulty))
			}
			if spawner.enemiesToSpawn > 30 {
				spawner.spawnTimer /= 2
			}
			spawner.enemiesToSpawn--
		}
	}

	world.enemySpawner = spawner
}

func enemyTypeToSpawn(world *World) *EnemyType {
	wave := world.enemySpawner.wave
	switch {
	case wave > 35 && rand.Float32() < 0.01:
		switch rand.Int() % 2 {
		case 0:
			return &enemyTypes.medicine
		case 1:
			return &enemyTypes.kogasa
		}
	case (world.difficulty == LUNATIC || wave > 20) && (rand.Float32() < 0.05 || (wave%6 == 0 && rand.Float32() < 0.2)):
		return &enemyTypes.purpleZombie
	case (wave%3 == 0 && rand.Float32() < 0.3) || rand.Float32() < 0.08:
		return &enemyTypes.smallZombie
	}
	return &enemyTypes.zombie
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
			updateRangedEnemy(world, id)
		}
	}
}

func updateRangedEnemy(world *World, id Entity) {
	world.shootTimer[id] -= dt
	if world.shootTimer[id] <= 0 {
		typ := world.enemy[id]
		world.shootTimer[id] = 1 - float32(world.difficulty)/10

		dir := util.Vector2Direction(world.position[id], world.position[world.player])
		vel := rl.Vector2Scale(dir, 100)
		newProjectile(world, world.position[id], vel, typ.projectileType)

		count := rand.Int()%3 + int(world.difficulty)
		for i := range count {
			vel = rl.Vector2Rotate(vel, math.Pi*2*float32(i)/float32(count))
			newProjectile(world, world.position[id], vel, typ.projectileType)
		}

		if world.enemySpawner.wave > 20 {
			vel = rl.Vector2Scale(vel, 0.8)
			for i := range count + 1 {
				vel = rl.Vector2Rotate(vel, math.Pi*2*float32(i-1)/float32(count))
				newProjectile(world, world.position[id], vel, typ.projectileType)
			}
		}
	}
}

func preEnemyDeath(world *World, id Entity) {
	switch world.enemy[id] {
	case &enemyTypes.purpleZombie:
		count := 2 + world.difficulty*3
		for i := range count {
			ratio := (float32(i) + 1) / float32(count)
			newProjectile(world, world.position[id], rl.Vector2Rotate(rl.Vector2{X: 80, Y: 0}, math.Pi*2*ratio), &projectileTypes.purpleBullet)
		}
	}

	world.playerData.mana += 1
	newDeathEffect(world, world.enemy[id].texture, world.position[id])

	stats.EnemiesKilled[world.difficulty]++
}
