package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func SpawnDoll(world *World) Entity {
	id := world.NewEntity()
	world.dollTag[id] = true
	world.position[id] = rl.Vector2Zero()
	world.velocity[id] = rl.Vector2Zero()
	return id
}
