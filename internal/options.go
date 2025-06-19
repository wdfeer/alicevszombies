package internal

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderOptions(world *World, origin rl.Vector2) {
	var maxTextWidth float32
	volumeTextSize := float32(rl.MeasureText("Volume", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	cursorTextSize := float32(rl.MeasureText("Cursor", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	maxTextWidth = max(volumeTextSize, cursorTextSize)

	buttonWidth := float32(480)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)

	buttonWidth -= maxTextWidth / 2
	soundVolume = raygui.SliderBar(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "", "Volume", soundVolume, 0, 1)

	origin.Y += buttonHeight + buttonSpacing
	raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(buttonWidth)/7)
	raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Cursor", &world.uistate.cursorType, 0, 1, false)

	origin.Y += buttonHeight + buttonSpacing
	buttonWidth += maxTextWidth
	if raygui.Button(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen") {
		rl.ToggleBorderlessWindowed()
	}
}
