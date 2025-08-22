package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type PixelParticle struct {
	timeleft   float32
	tint       rl.Color
	changeTint bool
	targetTint rl.Color
	alphaMode  uint8
}

func renderPixelParticles(world *World) {
	for id, eff := range world.pixelParticle {
		var alpha float32
		switch eff.alphaMode {
		case 0:
			alpha = min(eff.timeleft, 1)
		case 1:
			alpha = max(1-eff.timeleft, 0)
		case 2:
			alpha = 1
		}

		if options.Shadows {
			pos := rl.Vector2Add(world.position[id], rl.Vector2{X: 0.5, Y: 0.5})
			rl.DrawRectangleV(pos, rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(rl.Black, alpha))
		}

		tint := eff.tint
		if eff.changeTint {
			tint = rl.ColorLerp(eff.targetTint, tint, alpha)
		}

		rl.DrawRectangleV(world.position[id], rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(tint, alpha))
	}
}

func updatePixelParticles(world *World) {
	for id, eff := range world.pixelParticle {
		time := eff.timeleft - dt
		if time > 0 {
			world.pixelParticle[id] = PixelParticle{
				timeleft:   time,
				tint:       eff.tint,
				changeTint: eff.changeTint,
				targetTint: eff.targetTint,
				alphaMode:  eff.alphaMode,
			}
		} else {
			world.deleteEntity(id)
		}
	}
}
