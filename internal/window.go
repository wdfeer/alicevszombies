package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func InitWindowSettings() {
	monitor := rl.GetCurrentMonitor()
	fps := int32(rl.GetMonitorRefreshRate(monitor))
	rl.SetTargetFPS(fps)
	rl.SetExitKey(rl.KeyDelete)

	rl.SetWindowState(rl.FlagWindowResizable)
	rl.HideCursor()
}

func updateFullscreenToggleInput() {
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}
}
