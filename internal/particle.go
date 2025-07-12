package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type PixelParticle struct {
	timeleft     float32
	tint         rl.Color
	reverseAlpha bool
}

func renderPixelParticles(world *World) {
	for id, eff := range world.pixelParticle {
		var alpha float32
		if eff.reverseAlpha {
			alpha = max(1-eff.timeleft, 0)
		} else {
			alpha = min(eff.timeleft, 1)
		}

		if options.Shadows {
			pos := rl.Vector2Add(world.position[id], rl.Vector2{X: 0.5, Y: 0.5})
			rl.DrawRectangleV(pos, rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(rl.Black, alpha))
		}
		rl.DrawRectangleV(world.position[id], rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(eff.tint, alpha))
	}
}

func updatePixelParticles(world *World) {
	for id, eff := range world.pixelParticle {
		time := eff.timeleft - dt
		if time > 0 {
			world.pixelParticle[id] = PixelParticle{
				timeleft:     time,
				tint:         eff.tint,
				reverseAlpha: eff.reverseAlpha,
			}
		} else {
			world.deleteEntity(id)
		}
	}
}
