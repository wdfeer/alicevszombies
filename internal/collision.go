package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateCollisions(world *World) {
	playerRec := rl.Rectangle{
		X:      world.position[world.player].X,
		Y:      world.position[world.player].Y,
		Width:  world.size[world.player].X,
		Height: world.size[world.player].X}
	for id := range world.enemyTag {
		enemyRec := rl.Rectangle{
			X:      world.position[id].X,
			Y:      world.position[id].Y,
			Width:  world.size[id].X,
			Height: world.size[id].X}
		if rl.CheckCollisionRecs(playerRec, enemyRec) {
			damageWithCooldown(world, world.player, 1, id)
		}
	}
}
