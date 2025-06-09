package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(world *World) {
	camera := rl.Camera2D{}
	camera.Target = world.positions[world.player]
	camera.Offset = util.GetHalfScreen()
	camera.Zoom = 8

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)
	rl.BeginMode2D(camera)
	renderPlayer(world)
	renderEnemies(world)
	renderProjectiles(world)
	rl.EndMode2D()
}

func renderPlayer(world *World) {
	util.DrawTextureCentered(assets.Textures["player"], world.positions[world.player])
}

func renderEnemies(world *World) {
	// TODO
}

func renderProjectiles(world *World) {
	// TODO
}
