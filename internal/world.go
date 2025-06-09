package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity = uint32
type World struct {
	nextID        Entity
	player        Entity
	projectileTag map[Entity]bool
	enemyTag      map[Entity]bool
	positions     map[Entity]rl.Vector2
	velocities    map[Entity]rl.Vector2
}

func NewWorld() World {
	return World{
		nextID: 1,
	}
}

func (world *World) NewEntity() Entity {
	id := world.nextID
	world.nextID++
	return id
}

func (world *World) Update() {
	updateVelocity(world)

	rl.BeginDrawing()
	renderPlayer(world)
	renderEnemies(world)
	renderProjectiles(world)
	rl.EndDrawing()
}
