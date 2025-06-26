package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func renderTextures(world *World) {
	cameraRect := util.CenterRectangle(world.position[world.player], rl.Vector2Scale(util.ScreenSize(), 1/options.Zoom))
	items := make([]Entity, 0, len(world.texture))
	for id := range world.texture {
		pos, ok := world.position[id]
		if !ok {
			continue
		}

		if rl.CheckCollisionRecs(cameraRect, util.CenterRectangle(pos, rl.Vector2{X: 100, Y: 100})) {
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
		if texture == "knife" || texture == "magic_missile" {
			pos := util.CenterSomething(4, 4, pos)
			rl.DrawTextureEx(
				assets.textures[texture],
				pos,
				-rl.Vector2Angle(world.velocity[id], rl.Vector2{X: 1, Y: 0})*rl.Rad2deg,
				1,
				rl.White,
			)
		} else {
			util.DrawTextureCentered(assets.textures[texture], pos)
		}
	}
}
