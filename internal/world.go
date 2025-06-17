package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity = uint32
type World struct {
	paused       bool
	uistate      UIState
	nextID       Entity
	player       Entity
	playerData   PlayerData
	enemySpawner EnemySpawner
	targeting    map[Entity]Targeting
	doll         map[Entity]DollType
	enemy        map[Entity]*EnemyType
	projectile   map[Entity]Projectile
	position     map[Entity]rl.Vector2
	velocity     map[Entity]rl.Vector2
	drag         map[Entity]float32
	texture      map[Entity]string
	hp           map[Entity]HP
	combatText   map[Entity]CombatText
	size         map[Entity]rl.Vector2
	deathEffect  map[Entity]DeathEffectParticle
	animTimer    map[Entity]float32
	walkAnimated map[Entity]WalkAnimation
	rangedTimer  map[Entity]float32
}

func NewWorld() World {
	world := World{
		paused:       false,
		enemySpawner: newEnemySpawner(),
		targeting:    make(map[Entity]Targeting),
		doll:         make(map[Entity]DollType),
		enemy:        make(map[Entity]*EnemyType),
		position:     make(map[Entity]rl.Vector2),
		velocity:     make(map[Entity]rl.Vector2),
		drag:         make(map[Entity]float32),
		texture:      make(map[Entity]string),
		animTimer:    make(map[Entity]float32),
		hp:           make(map[Entity]HP),
		combatText:   make(map[Entity]CombatText),
		size:         make(map[Entity]rl.Vector2),
		deathEffect:  make(map[Entity]DeathEffectParticle),
		walkAnimated: make(map[Entity]WalkAnimation),
		projectile:   make(map[Entity]Projectile),
		rangedTimer:  make(map[Entity]float32),
		uistate:      UIState{},
	}

	newPlayer(&world)
	newDoll(&world, dollTypes.swordDoll)
	newDoll(&world, dollTypes.swordDoll)
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
		updateProjectiles(world)
		updateEnemySpawner(world)
		updateEnemies(world)
		updateTargetingMovement(world)
		updateDrag(world)
		updateVelocity(world)
		updateCollisions(world)

		updateDeathEffects(world)
		updateCombatText(world)
		updateAnimationData(world)
	}

	updateFullscreenToggleInput()
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
	} else if _, ok := world.enemy[entity]; ok {
		world.playerData.mana += 1
		newDeathEffect(world, "zombie", world.position[entity])
	}
	delete(world.targeting, entity)
	delete(world.doll, entity)
	delete(world.enemy, entity)
	delete(world.position, entity)
	delete(world.velocity, entity)
	delete(world.drag, entity)
	delete(world.texture, entity)
	delete(world.animTimer, entity)
	delete(world.hp, entity)
	delete(world.combatText, entity)
	delete(world.size, entity)
	delete(world.deathEffect, entity)
	delete(world.walkAnimated, entity)
	delete(world.projectile, entity)
	delete(world.rangedTimer, entity)
}
