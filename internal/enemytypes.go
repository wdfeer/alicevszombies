package internal

type EnemyType struct {
	texture      string
	acceleration float32
	baseHP       float32
	ranged       bool
}

var enemyTypes = struct {
	zombie    EnemyType
	redZombie EnemyType
	medicine  EnemyType
}{
	EnemyType{
		texture:      "zombie",
		acceleration: 700,
		baseHP:       3,
	},
	EnemyType{
		texture:      "red_zombie",
		acceleration: 800,
		baseHP:       2,
	},
	EnemyType{
		texture:      "medicine",
		acceleration: 730,
		baseHP:       50,
		ranged:       true,
	},
}
