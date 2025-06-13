package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerData struct {
	mana       float32
	dollDamage float32
}

func newPlayer(world *World) Entity {
	world.player = world.newEntity()
	world.position[world.player] = rl.Vector2Zero()
	world.velocity[world.player] = rl.Vector2Zero()
	world.drag[world.player] = 10
	world.hp[world.player] = HP{
		val:              10,
		immuneTime:       1.2,
		attackerCooldown: make(map[Entity]float32),
	}
	world.playerData = PlayerData{
		mana:       0,
		dollDamage: 1,
	}

	return world.player
}

const PLAYER_ACCELERATION = 3000

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

	delta := rl.Vector2Scale(dir, PLAYER_ACCELERATION*dt)
	world.velocity[world.player] = rl.Vector2Add(world.velocity[world.player], delta)
}
