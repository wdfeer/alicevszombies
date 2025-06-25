package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateCameraZoom() {
	if rl.IsKeyPressed(rl.KeyMinus) {
		options.Zoom = max(4, options.Zoom-1)
	}
	if rl.IsKeyPressed(rl.KeyEqual) {
		options.Zoom = min(12, options.Zoom+1)
	}
}

func createCamera(world *World) rl.Camera2D {
	camera := rl.Camera2D{
		Target: world.position[world.player],
		Offset: util.HalfScreenSize(),
		Zoom:   options.Zoom,
	}
	return camera
}
