package internal

import (
	"math"
	"math/rand"
)

type EnemySpawner struct {
	wave           uint32
	enemiesToSpawn uint32
	spawnTimer     float32
}

type SpawnData struct {
	boss    bool
	weight  float32
	minWave uint
	// Multiplied by the difficulty, this changes the min wave
	minWaveDiffMult int
}

func (data SpawnData) canSpawn(world *World) bool {
	wave := world.enemySpawner.wave

	minWave := max(int(data.minWave)+data.minWaveDiffMult*int(world.difficulty), 0)
	minWaveCondition := int(wave) >= minWave
	bossWaveCondition := (!data.boss || (wave%10 == 0 && wave > 0))
	onlyBossCondition := true

	if data.boss {
		for _, typ := range world.enemy {
			if typ.spawnData.boss {
				onlyBossCondition = false
			}
		}
	}

	return minWaveCondition && bossWaveCondition && onlyBossCondition
}

func updateEnemySpawner(world *World) {
	spawner := world.enemySpawner

	if spawner.enemiesToSpawn <= 0 {
		spawner.wave++
		spawner.enemiesToSpawn = 2 + spawner.wave*2
	}

	spawner.spawnTimer -= dt
	if world.status[world.player][Bleed] > 0 {
		spawner.spawnTimer -= dt * 2
	}

	if spawner.spawnTimer <= 0 {
		typ := enemyTypeToSpawn(world)
		newEnemy(world, typ)

		if typ.spawnData.boss {
			playSound("boss_spawn")
			spawner.enemiesToSpawn = 0
		} else {
			spawner.spawnTimer = 2 - min(1.4, float32(spawner.wave)/10)
			if spawner.enemiesToSpawn > 10 {
				spawner.spawnTimer /= max(2, float32(world.difficulty))
			}
			if spawner.enemiesToSpawn > 30 {
				spawner.spawnTimer /= float32(math.Exp2(float64(spawner.enemiesToSpawn) / 30))
			}
			spawner.enemiesToSpawn--
		}
	}

	world.enemySpawner = spawner
}

func enemyTypeToSpawn(world *World) *EnemyType {
	var valid []*EnemyType
	totalWeight := float32(0)
	for _, typ := range allEnemyTypes {
		if typ.spawnData.canSpawn(world) {
			valid = append(valid, typ)
			totalWeight += typ.spawnData.weight
		}
	}

	if len(valid) == 0 {
		println("WARNING: no valid enemy type to spawn found!")
		return &enemyTypes.zombie
	}

	val := rand.Float32() * totalWeight
	cumWeight := float32(0)
	for _, typ := range valid {
		cumWeight += typ.spawnData.weight
		if cumWeight > val {
			return typ
		}
	}

	println("WARNING: weighted enemy spawning failed!")
	return valid[rand.Int()%len(valid)]
}
