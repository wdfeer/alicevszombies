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
	cursorType int32
}

var options Options

func LoadOptions() {
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

	options = Options{
		fullscreen: true,
		volume:     1,
		cursorType: 0,
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
	newOptions := options

	var maxTextWidth float32
	volumeTextSize := float32(rl.MeasureText("Volume", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	cursorTextSize := float32(rl.MeasureText("Cursor", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	maxTextWidth = max(volumeTextSize, cursorTextSize)

	buttonWidth := float32(480)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)

	buttonWidth -= maxTextWidth / 2
	newOptions.volume = raygui.SliderBar(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "", "Volume", options.volume, 0, 1)

	origin.Y += buttonHeight + buttonSpacing
	raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(buttonWidth)/7)
	raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Cursor", &newOptions.cursorType, 0, 1, false)

	origin.Y += buttonHeight + buttonSpacing
	buttonWidth += maxTextWidth
	newOptions.fullscreen = raygui.Toggle(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen", options.fullscreen)

	if newOptions != options {
		options = newOptions
		saveOptions() // TODO: maybe run in background
	}
}
