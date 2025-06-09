package main

import (
	"alicevszombies/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1920, 1080, "alicevszombies")
	defer rl.CloseWindow()

	rl.SetTargetFPS(400)

	rl.SetWindowState(rl.FlagWindowResizable + rl.FlagBorderlessWindowedMode)

	world := internal.NewWorld()
	for !rl.WindowShouldClose() {
		world.Update()
	}
}
