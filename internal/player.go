package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerData struct {
	mana     float32
	upgrades map[*Upgrade]uint32
}

func (data *PlayerData) upgradeCount() uint32 {
	count := uint32(0)
	for _, v := range data.upgrades {
		count += v
	}
	return count
}

func newPlayer(world *World) Entity {
	world.player = world.newEntity()
	world.position[world.player] = rl.Vector2Zero()
	world.velocity[world.player] = rl.Vector2Zero()
	world.drag[world.player] = 10
	world.size[world.player] = rl.Vector2{X: 8, Y: 16}
	world.walkAnimated[world.player] = WalkAnimation{"player"}
	world.texture[world.player] = "player"

	return world.player
}

func updatePlayer(world *World) {
	dir := rl.Vector2Zero()

	inputMap := map[int32]rl.Vector2{
		rl.KeyW: {X: 0, Y: -1},
		rl.KeyA: {X: -1, Y: 0},
		rl.KeyS: {X: 0, Y: 1},
		rl.KeyD: {X: 1, Y: 0},
	}

	for k, v := range inputMap {
		if rl.IsKeyDown(k) {
			dir = rl.Vector2Add(dir, v)
		}
	}
	dir = rl.Vector2Normalize(dir)

	accel := float32(700 + world.playerData.upgrades[&MovementSpeed]*15)
	if world.status[world.player].slow > 0 {
		slowness := 10 + 10*world.difficulty
		accel -= float32(slowness)
	}
	delta := rl.Vector2Scale(dir, accel*dt)
	world.velocity[world.player] = rl.Vector2Add(world.velocity[world.player], delta)
}
