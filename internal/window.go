package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateFullscreenToggleInput() {
	if rl.IsKeyPressed(rl.KeyF) {
		if rl.IsWindowState(rl.FlagBorderlessWindowedMode) {
			rl.ClearWindowState(rl.FlagBorderlessWindowedMode)
		} else {
			rl.SetWindowState(rl.FlagBorderlessWindowedMode)
		}
	}
}
