package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DeathEffectParticle struct {
	timeleft float32
	color    rl.Color
}

type DeathEffectAsset struct {
	pixels map[rl.Vector2]rl.Color
	size   rl.Vector2
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
	assets.deathEffects[name] = DeathEffectAsset{pixels, rl.Vector2{X: float32(image.Width), Y: float32(image.Height)}}
}

func newDeathEffect(world *World, name string, center rl.Vector2) {
	position := rl.Vector2Subtract(center, rl.Vector2Scale(assets.deathEffects[name].size, CAMERA_ZOOM))
	for pixelpos, color := range assets.deathEffects[name].pixels {
		id := world.newEntity()
		world.position[id] = rl.Vector2Add(position, rl.Vector2Scale(pixelpos, CAMERA_ZOOM))
		world.velocity[id] = rl.Vector2Rotate(rl.Vector2{X: 0, Y: 1}, rand.Float32()-0.5)
		world.drag[id] = rand.Float32()/10 + 0.1
		world.deathEffect[id] = DeathEffectParticle{
			timeleft: 1,
			color:    color,
		}
	}
}

func updateDeathEffects(world *World) {
	for id, eff := range world.deathEffect {
		// TODO
	}
}

func renderDeathEffects(world *World) {
	// TODO
}
