package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func renderTextures(world *World, camera *rl.Camera2D) {
	cameraRect := util.CenterRectangle(camera.Target, rl.Vector2Scale(util.ScreenSize(), 1/float32(options.Zoom)))
	items := make([]Entity, 0)
	for id, textureName := range world.texture {
		pos, ok := world.position[id]
		if !ok {
			continue
		}

		if util.CheckCollisionRecs(cameraRect, util.CenterRectangle(pos, rl.Vector2{X: float32(assets.textures[textureName].Width), Y: float32(assets.textures[textureName].Height)})) {
			items = append(items, id)
		}
	}
	slices.Sort(items)

	renderNeededTextures(world, items)
}

func renderNeededTextures(world *World, ids []Entity) {
	for _, id := range ids {
		texture := world.texture[id]
		centeredPos := util.CenterSomething(float32(assets.textures[texture].Width), float32(assets.textures[texture].Height), world.position[id])
		rotation := float32(0)

		if texture == "knife" || texture == "magic_missile" || texture == "lightning_bolt" {
			rotation = -rl.Vector2Angle(world.velocity[id], rl.Vector2{X: 1, Y: 0}) * rl.Rad2deg
		}

		if options.Shadows {
			shadowOffset := rl.Vector2{X: 0.5, Y: 0.5}
			rl.DrawTextureEx(
				assets.textures[texture],
				rl.Vector2Add(centeredPos, shadowOffset),
				rotation,
				1,
				rl.Black,
			)
		}
		rl.DrawTextureEx(
			assets.textures[texture],
			centeredPos,
			rotation,
			1,
			rl.White,
		)
	}
}
