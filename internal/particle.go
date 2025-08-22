package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PixelParticle struct {
	timeleft     float32
	alphaMode    uint8
	tint         rl.Color
	changeTint   bool
	targetTint   rl.Color
	easingFactor float32
	initialPos   rl.Vector2
	targetPos    rl.Vector2
}

func renderPixelParticles(world *World) {
	if options.Shadows {
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

			pos := rl.Vector2Add(world.position[id], rl.Vector2{X: 0.5, Y: 0.5})
			rl.DrawRectangleV(pos, rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(rl.Black, alpha))
		}
	}

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

		tint := eff.tint
		if eff.changeTint {
			tint = rl.ColorLerp(tint, eff.targetTint, max(1-eff.timeleft, 0))
		}

		rl.DrawRectangleV(world.position[id], rl.Vector2{X: 1, Y: 1}, rl.ColorAlpha(tint, alpha))
	}
}

func updatePixelParticles(world *World) {
	for id, eff := range world.pixelParticle {
		time := eff.timeleft - dt
		if time > 0 {
			world.pixelParticle[id] = PixelParticle{
				timeleft:     time,
				tint:         eff.tint,
				changeTint:   eff.changeTint,
				targetTint:   eff.targetTint,
				alphaMode:    eff.alphaMode,
				easingFactor: eff.easingFactor,
				initialPos:   eff.initialPos,
				targetPos:    eff.targetPos,
			}
			if eff.targetPos.X != 0 && eff.targetPos.Y != 0 {
				world.position[id] = rl.Vector2Lerp(eff.initialPos, eff.targetPos, 1-min(eff.timeleft*eff.timeleft*eff.easingFactor, 1))
			}
		} else {
			world.deleteEntity(id)
		}
	}
}
