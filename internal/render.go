package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(world *World) {

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	renderGrid()

	rl.BeginMode2D(createCamera(world))
	renderPlayer(world)
	renderEnemies(world)
	renderProjectiles(world)
	rl.EndMode2D()
}

func renderGrid() {
	for i := -25; i <= 25; i++ {
		spacing := 1000 * int32(i)
		rl.DrawLine(-9999, spacing, 9999, spacing, rl.RayWhite)
		rl.DrawLine(spacing, -9999, spacing, 9999, rl.RayWhite)
	}
}

func createCamera(world *World) rl.Camera2D {
	camera := rl.Camera2D{
		Target: world.position[world.player],
		Offset: util.GetHalfScreen(),
		Zoom:   8,
	}
	camera.Target = world.position[world.player]
	camera.Offset = util.GetHalfScreen()
	camera.Zoom = 8
	return camera
}

func renderPlayer(world *World) {
	util.DrawTextureCentered(assets.Textures["player"], world.position[world.player])
}

func renderEnemies(world *World) {
	// TODO
}

func renderProjectiles(world *World) {
	// TODO
}
