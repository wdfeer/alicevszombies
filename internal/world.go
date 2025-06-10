package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity = uint32
type World struct {
	nextID        Entity
	player        Entity
	dollTag       map[Entity]bool
	projectileTag map[Entity]bool
	enemyTag      map[Entity]bool
	position      map[Entity]rl.Vector2
	velocity      map[Entity]rl.Vector2
	drag          map[Entity]float32
	texture       map[Entity]string
	animTimer     map[Entity]float32
}

func NewWorld() World {
	world := World{
		dollTag:       make(map[Entity]bool),
		projectileTag: make(map[Entity]bool),
		enemyTag:      make(map[Entity]bool),
		position:      make(map[Entity]rl.Vector2),
		velocity:      make(map[Entity]rl.Vector2),
		drag:          make(map[Entity]float32),
		texture:       make(map[Entity]string),
		animTimer:     make(map[Entity]float32),
	}

	newPlayer(&world)
	newDoll(&world)

	return world
}

func (world *World) Update() {
	updateInput(world)
	updateDrag(world)
	updateVelocity(world)

	updateAnimationData(world)
	render(world)
}

func (world *World) newEntity() Entity {
	id := world.nextID
	world.nextID++
	return id
}
