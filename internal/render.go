package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(world *World) {
	for i := range assets.renderTextures {
		if assets.renderTextures[i].Texture.Width != int32(rl.GetScreenWidth()) || assets.renderTextures[i].Texture.Height != int32(rl.GetScreenHeight()) {
			assets.renderTextures[i] = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
		}
	}

	rl.BeginTextureMode(assets.renderTextures[0])
	rl.ClearBackground(rl.Black)
	camera := createCamera(world)
	rl.BeginMode2D(camera)
	renderGrass(&camera)
	renderTextures(world, &camera)
	renderCombatText(world)
	renderPixelParticles(world)
	rl.EndMode2D()

	renderUI(world)
	rl.EndTextureMode()

	if options.Bloom {
		rl.BeginTextureMode(assets.renderTextures[1])
		rl.ClearBackground(rl.Black)
		rl.BeginShaderMode(assets.shaders["bloom"])
		rl.DrawTextureRec(assets.renderTextures[0].Texture,
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: -float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0}, rl.White)
		rl.EndShaderMode()
		rl.EndTextureMode()
	}

	rl.BeginDrawing()
	defer rl.EndDrawing()

	if options.Bloom && options.ChromaAbberation {
		rl.BeginShaderMode(assets.shaders["chroma_abberation"])
		rl.DrawTextureRec(assets.renderTextures[1].Texture,
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: -float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0}, rl.White)
		rl.EndShaderMode()
	} else if options.Bloom {
		rl.DrawTextureRec(assets.renderTextures[1].Texture,
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: -float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0}, rl.White)
	} else if options.ChromaAbberation {
		rl.BeginShaderMode(assets.shaders["chroma_abberation"])
		rl.DrawTextureRec(assets.renderTextures[0].Texture,
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: -float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0}, rl.White)
		rl.EndShaderMode()
	} else {
		rl.DrawTextureRec(assets.renderTextures[0].Texture,
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: -float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0}, rl.White)
	}
}

func renderGrass(camera *rl.Camera2D) {
	const GRASS_SIZE = 32
	origin := camera.Target
	origin.X -= util.ModF(origin.X, GRASS_SIZE)
	origin.Y -= util.ModF(origin.Y, GRASS_SIZE)
	for x := -400; x < 400; x += GRASS_SIZE {
		for y := -400; y < 400; y += GRASS_SIZE {
			pos := rl.Vector2{X: origin.X + float32(x), Y: origin.Y + float32(y)}
			rl.DrawTextureV(assets.textures["grass"], pos, rl.White)
		}
	}
}
