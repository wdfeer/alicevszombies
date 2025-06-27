package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func renderTextures(world *World) {
	cameraRect := util.CenterRectangle(world.position[world.player], rl.Vector2Scale(util.ScreenSize(), 1/options.Zoom))
	items := []Entity{}
	for id, textureName := range world.texture {
		pos, ok := world.position[id]
		if !ok {
			continue
		}

		if rl.CheckCollisionRecs(cameraRect, util.CenterRectangle(pos, rl.Vector2{X: float32(assets.textures[textureName].Width), Y: float32(assets.textures[textureName].Height)})) {
			items = append(items, id)
		}
	}
	slices.Sort(items)

	renderNeededTextures(world, items)
}

func renderNeededTextures(world *World, ids []Entity) {
	for _, id := range ids {
		texture := world.texture[id]
		pos := world.position[id]
		pos = util.CenterSomething(float32(assets.textures[texture].Width), float32(assets.textures[texture].Height), pos)
		rotation := float32(0)

		if texture == "knife" || texture == "magic_missile" {
			rotation = -rl.Vector2Angle(world.velocity[id], rl.Vector2{X: 1, Y: 0}) * rl.Rad2deg
		}

		shadowOffset := rl.Vector2{X: 0.5, Y: 0.5}
		pos = rl.Vector2Add(pos, shadowOffset)
		rl.DrawTextureEx(
			assets.textures[texture],
			pos,
			rotation,
			1,
			rl.Black,
		)
		pos = rl.Vector2Subtract(pos, shadowOffset)
		rl.DrawTextureEx(
			assets.textures[texture],
			pos,
			rotation,
			1,
			rl.White,
		)
	}
}
