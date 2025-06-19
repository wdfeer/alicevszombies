package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity = uint32
type World struct {
	paused       bool
	uistate      UIState
	difficulty   Difficulty
	nextID       Entity
	player       Entity
	playerData   PlayerData
	enemySpawner EnemySpawner
	targeting    map[Entity]Targeting
	doll         map[Entity]*DollType
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
	shootTimer   map[Entity]float32
}

func NewWorld() World {
	world := World{}
	world.Reset()
	return world
}

func (world *World) Reset() {
	{
		halfSize := util.HalfScreenSize()
		rl.SetMousePosition(int(halfSize.X), int(halfSize.Y))
	}

	world.paused = true
	world.enemySpawner = newEnemySpawner()
	world.targeting = make(map[Entity]Targeting)
	world.doll = make(map[Entity]*DollType)
	world.enemy = make(map[Entity]*EnemyType)
	world.position = make(map[Entity]rl.Vector2)
	world.velocity = make(map[Entity]rl.Vector2)
	world.drag = make(map[Entity]float32)
	world.texture = make(map[Entity]string)
	world.animTimer = make(map[Entity]float32)
	world.hp = make(map[Entity]HP)
	world.combatText = make(map[Entity]CombatText)
	world.size = make(map[Entity]rl.Vector2)
	world.deathEffect = make(map[Entity]DeathEffectParticle)
	world.walkAnimated = make(map[Entity]WalkAnimation)
	world.projectile = make(map[Entity]Projectile)
	world.shootTimer = make(map[Entity]float32)
	world.uistate = UIState{
		isMainMenu: true,
	}

	newPlayer(world)
	id := newDoll(world, &dollTypes.basicDoll)
	world.position[id] = rl.Vector2{X: -20, Y: 4}
	id = newDoll(world, &dollTypes.basicDoll)
	world.position[id] = rl.Vector2{X: 20, Y: -4}
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

	updateFullscreen()
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
		println("INFO: Player died!")
		onPlayerDeath(world)
	} else if _, ok := world.enemy[entity]; ok {
		preEnemyDeath(world, entity)
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
	delete(world.shootTimer, entity)
}
