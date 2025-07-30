package internal

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var uiScale float32 = 1
var textSize256 int64 = 256
var textSize64 int64 = 64
var textSize40 int64 = 40

func updateUIScale() {
	height := rl.GetScreenHeight()
	if height > 1080 {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 80)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 8)
		textSize256 = 256
		textSize64 = 64
		textSize40 = 40
		uiScale = 1
	} else if height > 720 {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 64)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 6)
		textSize256 = 160
		textSize64 = 48
		textSize40 = 32
		uiScale = 64. / 80
	} else {
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 4)
		textSize256 = 80
		textSize64 = 24
		textSize40 = 16
		uiScale = 32. / 80
	}
}
