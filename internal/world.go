package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity = uint32
type World struct {
	nextID        Entity
	player        Entity
	projectileTag map[Entity]bool
	enemyTag      map[Entity]bool
	position      map[Entity]rl.Vector2
	velocity      map[Entity]rl.Vector2
	drag          map[Entity]float32
}

func NewWorld() World {
	world := World{
		projectileTag: make(map[Entity]bool),
		enemyTag:      make(map[Entity]bool),
		position:      make(map[Entity]rl.Vector2),
		velocity:      make(map[Entity]rl.Vector2),
	}

	world.player = world.NewEntity()
	world.position[world.player] = rl.Vector2Zero()
	world.velocity[world.player] = rl.Vector2Zero()

	return world
}

func (world *World) NewEntity() Entity {
	id := world.nextID
	world.nextID++
	return id
}

func (world *World) Update() {
	updateInput(world)
	updateVelocity(world)

	render(world)
}
