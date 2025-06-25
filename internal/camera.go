package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var cameraZoom float32 = 8

func updateCameraZoom() {
	if rl.IsKeyPressed(rl.KeyMinus) {
		cameraZoom = max(4, cameraZoom-1)
	}
	if rl.IsKeyPressed(rl.KeyEqual) {
		cameraZoom = min(12, cameraZoom+1)
	}
}

func createCamera(world *World) rl.Camera2D {
	camera := rl.Camera2D{
		Target: world.position[world.player],
		Offset: util.HalfScreenSize(),
		Zoom:   cameraZoom,
	}
	return camera
}
