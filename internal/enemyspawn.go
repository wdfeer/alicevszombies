package internal

import "math/rand"

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

func (self SpawnData) canSpawn(world *World) bool {
	wave := world.enemySpawner.wave
	return wave >= uint32(int(self.minWave)+self.minWaveDiffMult*int(world.difficulty)) &&
		(!self.boss || (wave%10 == 0 && wave > 0))
}

func updateEnemySpawner(world *World) {
	spawner := world.enemySpawner

	if spawner.enemiesToSpawn <= 0 {
		spawner.wave++
		spawner.enemiesToSpawn = 2 + spawner.wave*2
	}

	spawner.spawnTimer = spawner.spawnTimer - dt
	if spawner.spawnTimer <= 0 {
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

	world.enemySpawner = spawner
}

func enemyTypeToSpawn(world *World) *EnemyType {
	valid := []*EnemyType{}
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
	cum_weight := float32(0)
	for _, typ := range valid {
		cum_weight += typ.spawnData.weight
		if cum_weight > val {
			return typ
		}
	}

	println("WARNING: weighted enemy spawning failed!")
	return valid[rand.Int()%len(valid)]
}
