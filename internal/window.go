package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateFullscreenToggleInput() {
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}
}
