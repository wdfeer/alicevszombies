package internal

import (
	"alicevszombies/internal/util"
	"fmt"
	"math/rand"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PauseMenu struct {
	options bool

	// Positions of dolls within the visual bouncy doll panel. Uses real world Entity ids
	dollPosition map[Entity]rl.Vector2
	// Velocities of dolls within the visual bouncy doll panel. Uses real world Entity ids
	dollVelocity map[Entity]rl.Vector2
}

func renderPauseMenu(world *World) {
	pMenu := &world.uistate.pauseMenu

	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	spacing := float32(40) * uiScale
	rectSize := rl.Vector2{X: 960 * uiScale, Y: 120 * uiScale}

	pos := rl.Vector2{X: spacing * 2, Y: spacing * 2}

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize256)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
	raygui.Label(util.RectangleV(pos, rectSize), "Paused")
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_CENTER))

	rectSize.X /= 2
	pos.Y += 256 * uiScale

	if raygui.Button(util.RectangleV(pos, rectSize), "Resume") {
		world.paused = false
	}

	pos.Y += rectSize.Y + spacing
	world.uistate.pauseMenu.options = raygui.Toggle(util.RectangleV(pos, rectSize), "Options", pMenu.options)

	if pMenu.options {
		pos := pos
		pos.X += rectSize.X + spacing*1.5
		pos.Y -= (rectSize.Y + spacing) * 1.5
		renderOptions(pos)
	}

	pos.Y += rectSize.Y + spacing
	if raygui.Button(util.RectangleV(pos, rectSize), "Main Menu") {
		world.Reset()
	}

	{ // Doll Counter
		oldAlignH := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT)
		oldAlignV := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT_VERTICAL)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, raygui.TEXT_ALIGN_RIGHT)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT_VERTICAL, raygui.TEXT_ALIGN_BOTTOM)
		rect := rl.Rectangle{X: float32(rl.GetScreenWidth()) - 400*uiScale, Y: float32(rl.GetScreenHeight()) - 128*uiScale, Width: 400 * uiScale, Height: 64 * uiScale}
		raygui.Label(rect, fmt.Sprintf("%d Dolls", len(world.doll)))
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, oldAlignH)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT_VERTICAL, oldAlignV)

		origin := rl.Vector2{X: 0, Y: float32(rl.GetScreenHeight()) - 64*uiScale}
		rect = rl.Rectangle{X: origin.X, Y: origin.Y, Width: float32(rl.GetScreenWidth()), Height: 64 * uiScale}
		raygui.Panel(rect, "")

		for id := range world.doll {
			if _, ok := pMenu.dollPosition[id]; !ok {
				pMenu.dollPosition[id] = rl.Vector2{X: float32(rl.GetScreenWidth()) * rand.Float32(), Y: 16 + 32*rand.Float32()}
				pMenu.dollVelocity[id] = rl.Vector2{X: (rand.Float32() - 0.5) * 200 * uiScale} // TODO: add vertical velocity and bounciness
			}

			pMenu.dollPosition[id] = rl.Vector2Add(pMenu.dollPosition[id], rl.Vector2Scale(pMenu.dollVelocity[id], dt))
			if pMenu.dollPosition[id].X >= float32(rl.GetScreenWidth()) {
				pMenu.dollPosition[id] = rl.Vector2{X: float32(rl.GetScreenWidth()), Y: pMenu.dollPosition[id].Y}
				pMenu.dollVelocity[id] = rl.Vector2{X: -pMenu.dollVelocity[id].X}
			} else if pMenu.dollPosition[id].X <= 0 {
				pMenu.dollPosition[id] = rl.Vector2{X: 0, Y: pMenu.dollPosition[id].Y}
				pMenu.dollVelocity[id] = rl.Vector2{X: -pMenu.dollVelocity[id].X}
			}

			util.DrawTextureCenteredScaled(assets.textures[world.texture[id]], rl.Vector2Add(origin, pMenu.dollPosition[id]), 3*uiScale)
		}
	}
}
