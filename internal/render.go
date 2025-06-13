package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(world *World) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(createCamera(world))
	renderGrid()
	renderTextures(world)
	renderCombatText(world)
	rl.EndMode2D()

	renderUI(world)
}

func renderTextures(world *World) {
	for id, texture := range world.texture {
		if pos, exists := world.position[id]; exists {
			util.DrawTextureCentered(assets.Textures[texture], pos)
		}
	}
}

func renderGrid() {
	for i := -25; i <= 25; i++ {
		spacing := 100 * int32(i)
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
