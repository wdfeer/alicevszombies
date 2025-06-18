package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type DollType struct {
	contactDamage  float32
	texture        string
	accel          float32
	projectileType *ProjectileType
	size           rl.Vector2
}

var dollTypes = struct {
	basicDoll    DollType
	lanceDoll    DollType
	knifeDoll    DollType
	magicianDoll DollType
	scytheDoll   DollType
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
		size:          rl.Vector2{X: 9, Y: 8},
	},
	scytheDoll: DollType{
		contactDamage: 3,
		texture:       "doll_scythe",
		accel:         650,
		size:          rl.Vector2{X: 13, Y: 8},
	},
	knifeDoll: DollType{
		contactDamage:  0,
		texture:        "doll_knife",
		accel:          400,
		projectileType: &projectileTypes.knife,
	},
	magicianDoll: DollType{
		contactDamage:  0,
		texture:        "doll_magician",
		accel:          550,
		projectileType: &projectileTypes.magicMissile,
	},
}
