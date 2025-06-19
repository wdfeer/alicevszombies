package internal

import (
	"alicevszombies/internal/util"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Options struct {
	fullscreen bool
	volume     float32
	cursorType uint8
}

var options = Options{
	fullscreen: true,
	volume:     1,
	cursorType: 0,
}

func loadOptions() {
	data, err := os.ReadFile("user/options.bin")
	if err == nil {
		if err = util.Deserialize(data, options); err == nil {
			println("INFO: Loaded options successfully!")
			return
		} else {
			println("ERROR: Failed deserializing options!")
		}
	} else {
		println("ERROR: Failed reading options file!")
	}

	saveOptions()
}

func saveOptions() {
	bytes, err := util.Serialize(options)
	if err != nil {
		println("ERROR: Failed serializing options!")
	}
	err = os.WriteFile("user/options.bin", bytes, 0644)
	if err != nil {
		println("ERROR: Failed writing options file!")
	}
}

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
