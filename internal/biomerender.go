package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderBiome(world *World, camera *rl.Camera2D) {
	renderGrass(world, camera)
	renderFoliage(world, camera)
}

func renderGrass(world *World, camera *rl.Camera2D) {
	const GRASS_SIZE = 32
	origin := camera.Target
	origin.X -= util.ModF(origin.X, GRASS_SIZE)
	origin.Y -= util.ModF(origin.Y, GRASS_SIZE)
	for x := -400; x < 400; x += GRASS_SIZE {
		for y := -400; y < 400; y += GRASS_SIZE {
			pos := rl.Vector2{X: origin.X + float32(x), Y: origin.Y + float32(y)}
			rl.DrawTextureV(assets.textures[world.worldgen.GetGrassTexture(pos)], pos, rl.White)
		}
	}
}

func renderFoliage(world *World, camera *rl.Camera2D) {
	start := rl.Vector2Subtract(camera.Target, camera.Offset)
	end := rl.Vector2Add(camera.Target, camera.Offset)
	foliage := world.worldgen.GetFoliageOnScreen(start, end)
	for pos, texture := range foliage {
		rl.DrawTextureV(assets.textures[texture], pos, rl.White)
	}
}
