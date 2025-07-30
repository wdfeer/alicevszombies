package internal

import (
	"alicevszombies/internal/util"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Options struct {
	optionsTab  uint8
	Fullscreen  bool
	MusicVolume float32
	SoundVolume float32
	CursorType  int32
	Zoom        int32
	Shadows     bool
	Bloom       bool
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
		Fullscreen:  true,
		MusicVolume: 1,
		CursorType:  0,
		Zoom:        8,
		Shadows:     true,
		Bloom:       true,
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

	buttonWidth := float32(616) * uiScale
	buttonHeight := float32(120) * uiScale
	buttonSpacing := float32(40) * uiScale

	buttonWidth -= maxTextWidth / 2

	{ // Tabs
		o := origin
		oldTextSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldTextSize/2)
		buttonWidth := 240. * uiScale
		buttonSpacing := buttonSpacing / 2
		buttonHeight := buttonHeight * 0.75
		if raygui.Toggle(rl.Rectangle{X: o.X, Y: o.Y, Width: buttonWidth, Height: buttonHeight}, "General", newOptions.optionsTab == 0) {
			newOptions.optionsTab = 0
		}
		o.X += buttonWidth + buttonSpacing
		if raygui.Toggle(rl.Rectangle{X: o.X, Y: o.Y, Width: buttonWidth, Height: buttonHeight}, "Audio", newOptions.optionsTab == 1) {
			newOptions.optionsTab = 1
		}
		o.X += buttonWidth + buttonSpacing
		if raygui.Toggle(rl.Rectangle{X: o.X, Y: o.Y, Width: buttonWidth, Height: buttonHeight}, "Graphics", newOptions.optionsTab == 2) {
			newOptions.optionsTab = 2
		}
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldTextSize)
		origin.Y += buttonHeight + buttonSpacing
	}

	switch newOptions.optionsTab {
	case 0:
		raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(buttonWidth)/7)
		raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Cursor", &newOptions.CursorType, 0, 1, false)
		origin.Y += buttonHeight + buttonSpacing
		raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Zoom", &newOptions.Zoom, MinZoom, MaxZoom, false)
	case 1:
		newOptions.MusicVolume = raygui.SliderBar(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "", "Music", options.MusicVolume, 0, 1)
		origin.Y += buttonHeight + buttonSpacing
		newOptions.SoundVolume = raygui.SliderBar(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "", "SFX", options.SoundVolume, 0, 1)
	case 2:
		buttonWidth += maxTextWidth
		newOptions.Fullscreen = raygui.Toggle(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Fullscreen", options.Fullscreen)
		origin.Y += buttonHeight + buttonSpacing
		newOptions.Shadows = raygui.Toggle(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Shadows", options.Shadows)
		origin.Y += buttonHeight + buttonSpacing
		newOptions.Bloom = raygui.Toggle(rl.Rectangle{X: origin.X, Y: origin.Y, Width: buttonWidth, Height: buttonHeight}, "Bloom", options.Bloom)
	}

	if newOptions != options {
		options = newOptions
		go saveOptions()
	}
}
