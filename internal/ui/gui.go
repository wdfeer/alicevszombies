package ui

import (
	"alicevszombies/internal/util"

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

// Draw texture at pos with scale, scaled with uiscale
func Texture(texture rl.Texture2D, pos rl.Vector2, scale float32) {
	rl.DrawTextureEx(texture, ScaleV(pos), 0, scale*UIScale, rl.White)
}

// Draw texture at center with scale, scaled with uiscale
func TextureC(texture rl.Texture2D, center rl.Vector2, scale float32) {
	pos := util.CenterSomethingV(rl.Vector2Scale(rl.Vector2{X: float32(texture.Width), Y: float32(texture.Height)}, scale*UIScale), center)
	rl.DrawTextureEx(texture, ScaleV(pos), 0, scale*UIScale, rl.White)
}
