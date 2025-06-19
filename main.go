package main

import (
	"alicevszombies/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "alicevszombies")
	defer rl.CloseWindow()

	internal.InitWindowSettings()

	internal.LoadAssets()
	defer internal.UnloadAssets()

	world := internal.NewWorld()
	for !rl.WindowShouldClose() {
		world.Update()
	}
}
