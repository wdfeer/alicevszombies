package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func renderTextures(world *World) {
	items := make([]Entity, 0, len(world.texture))
	for id := range world.texture {
		if _, ok := world.position[id]; ok {
			items = append(items, id)
		}
	}
	slices.Sort(items)

	for _, id := range items {
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
