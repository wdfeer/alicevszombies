package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderPauseMenu(world *World) {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	pos := util.HalfScreenSize()
	pos.Y *= 0.9

	rectSize := rl.Vector2{X: 900, Y: 120}
	spacing := float32(40)

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 256)
	raygui.Label(util.CenterRectangle(pos, rectSize), "Paused")
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)

	rectSize.X /= 2
	pos.Y += rectSize.Y + spacing*2

	if raygui.Button(util.CenterRectangle(pos, rectSize), "Resume") {
		world.paused = false
	}

	pos.Y += rectSize.Y + spacing

	if raygui.Button(util.CenterRectangle(pos, rectSize), "Quit") {
		rl.CloseWindow()
	}
}
