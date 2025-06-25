package ui

import rl "github.com/gen2brain/raylib-go/raylib"

var UIScale float32 = 1

func UpdateUIScale() {
	UIScale = float32(rl.GetScreenHeight()) / 1440
}

func ScaleV(vec rl.Vector2) rl.Vector2 {
	return rl.Vector2Scale(vec, UIScale)
}

func ScaleR(rect rl.Rectangle) rl.Rectangle {
	return rl.Rectangle{X: rect.X * UIScale, Y: rect.Y * UIScale, Width: rect.Width * UIScale, Height: rect.Height * UIScale}
}
