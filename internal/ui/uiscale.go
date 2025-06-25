package ui

import rl "github.com/gen2brain/raylib-go/raylib"

var UIScale float32 = 1

func UpdateUIScale() {
	UIScale = float32(rl.GetScreenHeight()) / 1440
}
