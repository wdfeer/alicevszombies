package internal

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var uiScale float32 = 1

func updateUIScale() {
	height := rl.GetScreenHeight()
	if height > 1080 {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 80)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 8)
		uiScale = 1
	} else if height > 720 {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 64)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 6)
		uiScale = 64. / 80
	} else {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 4)
		uiScale = 32. / 80
	}
}
