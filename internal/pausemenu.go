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

	pos := util.HalfScreenSize()
	pos.Y *= 0.5

	rectSize := rl.Vector2{X: 960, Y: 120}
	spacing := float32(40)

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 256)
	raygui.Label(util.CenterRectangle(pos, rectSize), "Paused")
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)

	rectSize.X /= 2
	pos = util.HalfScreenSize()

	if raygui.Button(util.CenterRectangle(pos, rectSize), "Resume") {
		world.paused = false
	}

	pos.Y += rectSize.Y + spacing
	world.uistate.pauseMenu.options = raygui.Toggle(util.CenterRectangle(pos, rectSize), "Options", world.uistate.pauseMenu.options)

	pos.Y += rectSize.Y + spacing
	if raygui.Button(util.CenterRectangle(pos, rectSize), "Main Menu") {
		world.Reset()
	}

	if world.uistate.pauseMenu.options {
		pos = util.HalfScreenSize()
		pos.X += rectSize.X * 2 / 3
		pos.Y -= rectSize.Y / 2
		renderOptions(pos)
	}
}
