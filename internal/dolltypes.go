package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type DollType struct {
	contactDamage float32
	texture       string
	accel         float32
	size          rl.Vector2
	shootPattern  ShootPattern
}

var dollTypes = struct {
	basicDoll       DollType
	lanceDoll       DollType
	knifeDoll       DollType
	magicianDoll    DollType
	scytheDoll      DollType
	destructionDoll DollType
}{
	basicDoll: DollType{
		contactDamage: 1,
		texture:       "doll_sword",
		accel:         500,
		size:          rl.Vector2{X: 8, Y: 8},
	},
	lanceDoll: DollType{
		contactDamage: 2,
		texture:       "doll_lance",
		accel:         500,
		size:          rl.Vector2{X: 11, Y: 8},
	},
	scytheDoll: DollType{
		contactDamage: 3,
		texture:       "doll_scythe",
		accel:         650,
		size:          rl.Vector2{X: 20, Y: 8},
	},
	knifeDoll: DollType{
		contactDamage: 0,
		texture:       "doll_knife",
		accel:         400,
		shootPattern: ShootPattern{
			projectile: &projectileTypes.knife,
			cooldown:   1,
			typ:        Direct,
			velocity:   200,
		},
	},
	magicianDoll: DollType{
		contactDamage: 0,
		texture:       "doll_magician",
		accel:         550,
		shootPattern: ShootPattern{
			projectile: &projectileTypes.magicMissile,
			cooldown:   1,
			typ:        Direct,
			velocity:   200,
		},
	},
	destructionDoll: DollType{
		contactDamage: 5,
		texture:       "doll_destruction",
		accel:         400,
		size:          rl.Vector2{X: 18, Y: 10},
		shootPattern: ShootPattern{
			projectile: &projectileTypes.redBullet,
			cooldown:   1,
			typ:        Circle,
			count:      5,
			velocity:   200,
		},
	},
}
