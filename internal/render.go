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
	renderGrass(world)
	renderTextures(world)
	renderCombatText(world)
	renderDeathEffects(world)
	rl.EndMode2D()

	renderUI(world)
}

func renderGrass(world *World) {
	const GRASS_SIZE = 32
	origin := rl.Vector2Subtract(world.position[world.player], rl.Vector2Scale(util.HalfScreenSize(), float32(1)/CAMERA_ZOOM))
	origin.X = origin.X - util.ModF(origin.X, GRASS_SIZE)
	origin.Y = origin.Y - util.ModF(origin.Y, GRASS_SIZE)
	for x := -400; x < 400; x += GRASS_SIZE {
		for y := -400; y < 400; y += GRASS_SIZE {
			pos := rl.Vector2{X: origin.X + float32(x), Y: origin.Y + float32(y)}
			rl.DrawTextureV(assets.textures["grass"], pos, rl.White)
		}
	}
}

const CAMERA_ZOOM = 8

func createCamera(world *World) rl.Camera2D {
	camera := rl.Camera2D{
		Target: world.position[world.player],
		Offset: util.HalfScreenSize(),
		Zoom:   CAMERA_ZOOM,
	}
	return camera
}
