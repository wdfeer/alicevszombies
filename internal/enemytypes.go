package internal

type EnemyType struct {
	texture      string
	acceleration float32
	baseHP       float32
}

var enemyTypes = struct {
	zombie   EnemyType
	medicine EnemyType
}{
	EnemyType{
		texture:      "zombie",
		acceleration: 700,
		baseHP:       3,
	},
	EnemyType{
		texture:      "medicine",
		acceleration: 730,
		baseHP:       50,
	},
}
