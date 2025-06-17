package internal

type EnemyType struct {
}

var enemyTypes = struct {
	zombie   EnemyType
	medicine EnemyType
}{
	EnemyType{},
	EnemyType{},
}
