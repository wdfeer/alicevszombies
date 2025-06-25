package internal

import (
	"alicevszombies/internal/util"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Options struct {
	Fullscreen bool
	Volume     float32
	CursorType int32
	Zoom       float32
}

var options Options

func loadOptions() {
	data, err := os.ReadFile("user/options.bin")
	if err == nil {
		if err = util.Deserialize(data, &options); err == nil {
			println("INFO: Loaded options successfully!")
			return
		} else {
			println("ERROR: Failed deserializing options!")
		}
	} else {
		println("ERROR: Failed reading options file!")
	}

	println("WARNING: Creating default options file...")

	options = Options{
		Fullscreen: true,
		Volume:     1,
		CursorType: 0,
		Zoom:       8,
	}

	go saveOptions()
}

func saveOptions() {
	bytes, err := util.Serialize(&options)
	if err != nil {
		println("ERROR: Failed serializing options!")
		return
	}

	if _, err = os.Stat("user"); err != nil {
		err = os.Mkdir("user", 0755)
		if err != nil {
			println("ERROR: Failed creating \"user\" directory!")
			return
		}
	}

	err = os.WriteFile("user/options.bin", bytes, 0644)
	if err != nil {
		println("ERROR: Failed writing options file!")
		return
	}
	println("INFO: Options saved!")
}

func renderOptions(origin rl.Vector2) {
	newOptions := options

	var maxTextWidth float32
	volumeTextSize := float32(rl.MeasureText("Volume", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	cursorTextSize := float32(rl.MeasureText("Cursor", int32(raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE))))
	maxTextWidth = max(volumeTextSize, cursorTextSize)

	buttonWidth := float32(480)
	buttonHeight := float32(120)
	buttonSpacing := float32(40)

	buttonWidth -= maxTextWidth / 2
	newOptions.Volume = raygui.SliderBar(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "", "Volume", options.Volume, 0, 1)

	origin.Y += buttonHeight + buttonSpacing
	raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(buttonWidth)/7)
	raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Cursor", &newOptions.CursorType, 0, 1, false)

	origin.Y += buttonHeight + buttonSpacing
	buttonWidth += maxTextWidth
	newOptions.Fullscreen = raygui.Toggle(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen", options.Fullscreen)

	if newOptions != options {
		options = newOptions
		go saveOptions()
	}
}
