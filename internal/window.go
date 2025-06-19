package internal

import (
	"alicevszombies/internal/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitWindowSettings() {
	monitor := rl.GetCurrentMonitor()
	fps := int32(rl.GetMonitorRefreshRate(monitor))
	rl.SetTargetFPS(fps)
	rl.SetExitKey(rl.KeyDelete)

	size := rl.Vector2Scale(util.ScreenSize(), 0.8)
	pos := util.CenterSomethingV(size, util.HalfScreenSize())
	rl.SetWindowPosition(int(pos.X), int(pos.Y))
	rl.SetWindowSize(int(size.X), int(size.Y))
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.HideCursor()
}

func updateFullscreen() {
	if rl.IsKeyPressed(rl.KeyF) {
		options.Fullscreen = !options.Fullscreen
		saveOptions()
	}

	isFullscreen := rl.IsWindowState(rl.FlagBorderlessWindowedMode)
	if (options.Fullscreen && !isFullscreen) || (!options.Fullscreen && isFullscreen) {
		rl.ToggleBorderlessWindowed()
	}
}
