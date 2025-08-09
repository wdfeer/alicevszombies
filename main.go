package main

import (
	"alicevszombies/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	internal.InitWindow()
	defer rl.CloseWindow()

	internal.LoadUserData()

	internal.LoadAssets()
	defer internal.UnloadAssets()

	world := internal.NewWorld()

	for !rl.WindowShouldClose() {
		world.Update()
	}

	internal.SaveRun(&world)

	// FIXME: this doesn't seem to actually call the function
	internal.SaveUserData()
}
