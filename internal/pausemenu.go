package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PauseMenu struct {
	options bool
}

func renderPauseMenu(world *World) {
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
	world.uistate.pauseMenu.options = raygui.Toggle(util.RectangleV(pos, rectSize), "Options", world.uistate.pauseMenu.options)

	if world.uistate.pauseMenu.options {
		pos := pos
		pos.X += rectSize.X + spacing*1.5
		pos.Y -= (rectSize.Y + spacing) * 1.5
		renderOptions(pos)
	}

	pos.Y += rectSize.Y + spacing
	if raygui.Button(util.RectangleV(pos, rectSize), "Main Menu") {
		world.Reset()
	}
}
