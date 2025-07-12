package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureBreakdown struct {
	pixels map[rl.Vector2]rl.Color
	size   rl.Vector2
}

func loadBreakdown(name string) {
	pixels := make(map[rl.Vector2]rl.Color)
	texture, exists := assets.textures[name]
	if !exists {
		panic("Texture with name \"" + name + "\" not found!")
	}

	image := rl.LoadImageFromTexture(texture)
	colors := rl.LoadImageColors(image)
	for x := range image.Width {
		for y := range image.Height {
			clr := colors[y*image.Width+x]
			if clr.A > 0 {
				pixels[rl.Vector2{X: float32(x), Y: float32(y)}] = clr
			}
		}
	}
	assets.breakdowns[name] = TextureBreakdown{pixels, rl.Vector2{X: float32(image.Width), Y: float32(image.Height)}}
}

// Spawns pixel particles based on a TextureBreakdown
// Returns all the PixelParticle ids
func newBreakdown(world *World, name string, center rl.Vector2) []Entity {
	position := rl.Vector2Subtract(center, rl.Vector2Scale(assets.breakdowns[name].size, 0.5))
	entities := make([]Entity, len(assets.breakdowns[name].pixels))
	i := 0
	for pixelpos, color := range assets.breakdowns[name].pixels {
		id := world.newEntity()
		entities[i] = id
		i++

		world.position[id] = rl.Vector2Add(position, pixelpos)
		world.pixelParticle[id] = PixelParticle{
			timeleft: 1,
			tint:     color,
		}
	}
	return entities
}
