package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type DeathEffect struct {
	pixels map[rl.Vector2]rl.Color
}

func loadDeathEffect(name string) {
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
	assets.deathEffects[name] = DeathEffect{pixels}
}
