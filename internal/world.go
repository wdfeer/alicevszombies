package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity = uint32
type World struct {
	paused       bool
	nextID       Entity
	player       Entity
	playerData   PlayerData
	enemySpawner EnemySpawner
	targeting    map[Entity]Targeting
	dollTag      map[Entity]bool
	enemyTag     map[Entity]bool
	position     map[Entity]rl.Vector2
	velocity     map[Entity]rl.Vector2
	drag         map[Entity]float32
	texture      map[Entity]string
	animTimer    map[Entity]float32
	hp           map[Entity]HP
	combatText   map[Entity]CombatText
}

func NewWorld() World {
	world := World{
		paused:       false,
		enemySpawner: newEnemySpawner(),
		targeting:    make(map[Entity]Targeting),
		dollTag:      make(map[Entity]bool),
		enemyTag:     make(map[Entity]bool),
		position:     make(map[Entity]rl.Vector2),
		velocity:     make(map[Entity]rl.Vector2),
		drag:         make(map[Entity]float32),
		texture:      make(map[Entity]string),
		animTimer:    make(map[Entity]float32),
		hp:           make(map[Entity]HP),
		combatText:   make(map[Entity]CombatText),
	}

	newPlayer(&world)
	newDoll(&world)
	newDoll(&world)
	newEnemy(&world)

	return world
}

var dt float32

func (world *World) Update() {
	dt = rl.Clamp(rl.GetFrameTime(), 0.002, 0.05)

	if !world.paused {
		updateHP(world)
		updatePlayer(world)
		updateSpells(world)
		updateDolls(world)
		updateEnemySpawner(world)
		updateEnemies(world)
		updateTargetingMovement(world)
		updateDrag(world)
		updateVelocity(world)
		updateCombatText(world)
		updateAnimationData(world)
	}

	updateUI(world)
	render(world)
}

func (world *World) newEntity() Entity {
	id := world.nextID
	world.nextID++
	return id
}

func (world *World) deleteEntity(entity Entity) {
	if world.player == entity {
		println("Player died! Closing the game.")
		rl.CloseWindow()
	} else if world.enemyTag[entity] {
		world.playerData.mana += 1
	}
	delete(world.targeting, entity)
	delete(world.dollTag, entity)
	delete(world.enemyTag, entity)
	delete(world.position, entity)
	delete(world.velocity, entity)
	delete(world.drag, entity)
	delete(world.texture, entity)
	delete(world.animTimer, entity)
	delete(world.hp, entity)
	delete(world.combatText, entity)
}
