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
	renderGrass()
	renderTextures(world)
	renderCombatText(world)
	renderDeathEffects(world)
	rl.EndMode2D()

	renderUI(world)
}

func renderTextures(world *World) {
	for id, texture := range world.texture {
		if pos, exists := world.position[id]; exists {
			util.DrawTextureCentered(assets.textures[texture], pos)
		}
	}
}

func renderGrass() {
	// TODO: only draw visible grass
	for x := int32(-500); x < 500; x += 16 {
		for y := int32(-500); y < 500; y += 16 {
			rl.DrawTexture(assets.textures["grass"], x, y, rl.White)
		}
	}
}

const CAMERA_ZOOM = 8

func createCamera(world *World) rl.Camera2D {
	camera := rl.Camera2D{
		Target: world.position[world.player],
		Offset: util.GetHalfScreen(),
		Zoom:   CAMERA_ZOOM,
	}
	return camera
}
