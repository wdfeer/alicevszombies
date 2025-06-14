package main

import (
	"alicevszombies/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	monitor := rl.GetCurrentMonitor()

	rl.InitWindow(int32(rl.GetMonitorWidth(monitor)), int32(rl.GetMonitorHeight(monitor)), "alicevszombies")
	defer rl.CloseWindow()

	fps := int32(rl.GetMonitorRefreshRate(monitor))
	rl.SetTargetFPS(fps)
	rl.SetExitKey(rl.KeyDelete)

	rl.SetWindowState(rl.FlagWindowResizable + rl.FlagBorderlessWindowedMode)
	rl.HideCursor()

	internal.LoadAssets()
	world := internal.NewWorld()
	for !rl.WindowShouldClose() {
		world.Update()
	}
}
