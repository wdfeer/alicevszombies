package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	typ      ProjectileType
	timeLeft float32
}

func newProjectile(world *World, pos rl.Vector2, vel rl.Vector2, typ ProjectileType) Entity {
	id := world.newEntity()
	world.position[id] = pos
	world.velocity[id] = vel
	//TODO
	return id
}

type ProjectileType = uint8 // TODO: make struct

const (
	KNIFE_PROJECTILE ProjectileType = iota
)
