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
	deathExplode   DeathExplode
}

type DeathExplode struct {
	active         bool
	projectileType *ProjectileType
	countFlat      uint
	countDiffMult  uint
}

func (self DeathExplode) getProjectileCount(world *World) uint {
	return self.countFlat + self.countDiffMult*uint(world.difficulty)
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
	},
	EnemyType{
		texture:      "small_zombie",
		acceleration: 740,
		baseHP:       1,
		size:         rl.Vector2{X: 4, Y: 8},
	},
	EnemyType{
		texture:      "purple_zombie",
		acceleration: 700,
		baseHP:       2,
		size:         rl.Vector2{X: 8, Y: 16},
		deathExplode: DeathExplode{
			active:         true,
			projectileType: &projectileTypes.purpleBullet,
			countFlat:      4,
			countDiffMult:  1,
		},
	},
	EnemyType{
		texture:      "blue_zombie",
		acceleration: 700,
		baseHP:       4,
		size:         rl.Vector2{X: 8, Y: 16},
		deathExplode: DeathExplode{
			active:         true,
			projectileType: &projectileTypes.blueBullet,
			countFlat:      2,
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
	},
	EnemyType{
		texture:        "kogasa",
		acceleration:   715,
		baseHP:         70,
		ranged:         true,
		size:           rl.Vector2{X: 8, Y: 16},
		flippable:      true,
		projectileType: &projectileTypes.blueBullet,
	},
}
