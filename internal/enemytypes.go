package internal

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyType struct {
	texture        string
	acceleration   float32
	drag           float32
	baseHP         float32
	projectileType *ProjectileType
	size           rl.Vector2
	disableWalking bool
	flippable      bool
	targetingType  EnemyTargetingType
	spawnData      SpawnData
	shootPattern   ShootPattern
	deathExplode   DeathExplode
	// don't collide with other enemies
	flying bool
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
	neriumGirl   EnemyType
	zombieFairy  EnemyType
	medicine     EnemyType
	kogasa       EnemyType
}{
	EnemyType{
		texture:      "zombie",
		acceleration: 680,
		drag:         10,
		baseHP:       3,
		size:         rl.Vector2{X: 8, Y: 16},
		spawnData: SpawnData{
			weight: 1,
		},
		targetingType: DirectMelee,
	},
	EnemyType{
		texture:      "small_zombie",
		acceleration: 740,
		drag:         10,
		baseHP:       1,
		size:         rl.Vector2{X: 4, Y: 8},
		spawnData: SpawnData{
			weight:  0.2,
			minWave: 4,
		},
		targetingType: DirectMelee,
	},
	EnemyType{
		texture:      "purple_zombie",
		acceleration: 700,
		drag:         10,
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
			countDiffMult:  2,
		},
		targetingType: DirectMelee,
	},
	EnemyType{
		texture:      "blue_zombie",
		acceleration: 690,
		drag:         10,
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
			countFlat:      4,
			countDiffMult:  2,
		},
		targetingType: DirectMelee,
	},
	EnemyType{
		texture:      "nerium_girl",
		acceleration: 725,
		drag:         10,
		baseHP:       9,
		size:         rl.Vector2{X: 8, Y: 16},
		spawnData: SpawnData{
			weight:          0.02,
			minWave:         35,
			minWaveDiffMult: -3,
		},
		shootPattern: ShootPattern{
			projectile:              &projectileTypes.purpleBullet,
			cooldown:                2.5,
			typ:                     SpreadShoot,
			count:                   2,
			countExtraPerDifficulty: 1.25,
			spread:                  math.Pi / 3,
			velocity:                80,
		},
		targetingType: Ranged,
	},
	EnemyType{
		texture:      "zombie_fairy",
		acceleration: 600,
		drag:         5,
		baseHP:       2,
		size:         rl.Vector2{X: 7, Y: 8},
		spawnData: SpawnData{
			weight:          0.06,
			minWave:         29,
			minWaveDiffMult: -3,
		},
		targetingType:  CirclingMelee,
		disableWalking: true,
		flying:         true,
	},
	EnemyType{
		texture:        "medicine",
		acceleration:   730,
		drag:           10,
		baseHP:         50,
		size:           rl.Vector2{X: 8, Y: 16},
		projectileType: &projectileTypes.purpleBullet,
		spawnData: SpawnData{
			boss:   true,
			weight: 1,
		},
		shootPattern: ShootPattern{
			projectile:              &projectileTypes.purpleBullet,
			cooldown:                1,
			typ:                     CircleShoot,
			count:                   4,
			countExtraPerDifficulty: 1,
			countExtraPerWave:       0.05,
			velocity:                100,
		},
		targetingType: Ranged,
	},
	EnemyType{
		texture:        "kogasa",
		acceleration:   715,
		drag:           10,
		baseHP:         70,
		size:           rl.Vector2{X: 8, Y: 16},
		flippable:      true,
		projectileType: &projectileTypes.blueBullet,
		spawnData: SpawnData{
			boss:   true,
			weight: 1,
		},
		shootPattern: ShootPattern{
			projectile:              &projectileTypes.blueBullet,
			cooldown:                1,
			typ:                     CircleShoot,
			count:                   4,
			countExtraPerDifficulty: 1,
			countExtraPerWave:       0.05,
			velocity:                100,
		},
		targetingType: LeadingMelee,
	},
}

var allEnemyTypes = []*EnemyType{
	&enemyTypes.zombie,
	&enemyTypes.smallZombie,
	&enemyTypes.purpleZombie,
	&enemyTypes.blueZombie,
	&enemyTypes.neriumGirl,
	&enemyTypes.zombieFairy,
	&enemyTypes.medicine,
	&enemyTypes.kogasa,
}
