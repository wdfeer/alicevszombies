package internal

import (
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
	renderBiome(world, &camera)
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
		offsetLocation := rl.GetShaderLocation(assets.shaders["chroma_abberation"], "offset")
		rl.SetShaderValue(assets.shaders["chroma_abberation"], offsetLocation, []float32{0.0006 * 2560 / float32(rl.GetScreenWidth())}, rl.ShaderUniformFloat)
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
		offsetLocation := rl.GetShaderLocation(assets.shaders["chroma_abberation"], "offset")
		rl.SetShaderValue(assets.shaders["chroma_abberation"], offsetLocation, []float32{0.0006 * 2560 / float32(rl.GetScreenWidth())}, rl.ShaderUniformFloat)
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
