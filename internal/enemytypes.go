package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type EnemyType struct {
	texture      string
	acceleration float32
	baseHP       float32
	ranged       bool
	size         rl.Vector2
}

var enemyTypes = struct {
	zombie      EnemyType
	smallZombie EnemyType
	redZombie   EnemyType
	medicine    EnemyType
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
		texture:      "red_zombie",
		acceleration: 700,
		baseHP:       2,
		size:         rl.Vector2{X: 8, Y: 16},
	},
	EnemyType{
		texture:      "medicine",
		acceleration: 730,
		baseHP:       50,
		ranged:       true,
		size:         rl.Vector2{X: 8, Y: 16},
	},
}
