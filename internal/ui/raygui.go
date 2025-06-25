package ui

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Raygui label with scaled position and size
func Label(rect rl.Rectangle, text string) {
	raygui.Label(ScaleR(rect), text)
}

// Raygui label button with scaled position and size
func LabelButton(rect rl.Rectangle, text string) bool {
	return raygui.LabelButton(ScaleR(rect), text)
}

// Raygui button with scaled position and size
func Button(rect rl.Rectangle, text string) bool {
	return raygui.Button(ScaleR(rect), text)
}

// Raygui toggle with scaled position and size
func Toggle(rect rl.Rectangle, text string, active bool) bool {
	return raygui.Toggle(ScaleR(rect), text, active)
}

// Raygui spinner with scaled position and size
func Spinner(rect rl.Rectangle, text string, value *int32, min int, max int) bool {
	return raygui.Spinner(ScaleR(rect), text, value, min, max, false)
}
