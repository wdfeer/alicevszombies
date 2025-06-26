package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type EnemyType struct {
	texture        string
	acceleration   float32
	baseHP         float32
	ranged         bool
	projectileType *ProjectileType
	size           rl.Vector2
	flippable      bool
	spawnData      SpawnData
	deathExplode   DeathExplode
}

type DeathExplode struct {
	active         bool
	projectileType *ProjectileType
	// Flat projectile count summand
	countFlat uint
	// Projectile count summand that is to be multiplied by difficulty
	countDiffMult uint
}

func (exploding DeathExplode) getProjectileCount(world *World) uint {
	return exploding.countFlat + exploding.countDiffMult*uint(world.difficulty)
}

var enemyTypes = struct {
	zombie       EnemyType
	smallZombie  EnemyType
	purpleZombie EnemyType
	blueZombie   EnemyType
	medicine     EnemyType
	kogasa       EnemyType
}{
	EnemyType{
		texture:      "zombie",
		acceleration: 680,
		baseHP:       3,
		size:         rl.Vector2{X: 8, Y: 16},
		spawnData: SpawnData{
			weight: 1,
		},
	},
	EnemyType{
		texture:      "small_zombie",
		acceleration: 740,
		baseHP:       1,
		size:         rl.Vector2{X: 4, Y: 8},
		spawnData: SpawnData{
			weight:  0.2,
			minWave: 4,
		},
	},
	EnemyType{
		texture:      "purple_zombie",
		acceleration: 700,
		baseHP:       2,
		size:         rl.Vector2{X: 8, Y: 16},
		spawnData: SpawnData{
			weight:          0.05,
			minWave:         26,
			minWaveDiffMult: -6,
		},
		deathExplode: DeathExplode{
			active:         true,
			projectileType: &projectileTypes.purpleBullet,
			countFlat:      4,
			countDiffMult:  1,
		},
	},
	EnemyType{
		texture:      "blue_zombie",
		acceleration: 690,
		baseHP:       4,
		size:         rl.Vector2{X: 8, Y: 16},
		spawnData: SpawnData{
			weight:          0.05,
			minWave:         26,
			minWaveDiffMult: -6,
		},
		deathExplode: DeathExplode{
			active:         true,
			projectileType: &projectileTypes.blueBullet,
			countFlat:      3,
			countDiffMult:  2,
		},
	},
	EnemyType{
		texture:        "medicine",
		acceleration:   730,
		baseHP:         50,
		ranged:         true,
		size:           rl.Vector2{X: 8, Y: 16},
		projectileType: &projectileTypes.purpleBullet,
		spawnData: SpawnData{
			boss:   true,
			weight: 1,
		},
	},
	EnemyType{
		texture:        "kogasa",
		acceleration:   715,
		baseHP:         70,
		ranged:         true,
		size:           rl.Vector2{X: 8, Y: 16},
		flippable:      true,
		projectileType: &projectileTypes.blueBullet,
		spawnData: SpawnData{
			boss:   true,
			weight: 1,
		},
	},
}

var allEnemyTypes = []*EnemyType{
	&enemyTypes.zombie,
	&enemyTypes.smallZombie,
	&enemyTypes.purpleZombie,
	&enemyTypes.blueZombie,
	&enemyTypes.medicine,
	&enemyTypes.kogasa,
}
