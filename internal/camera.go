package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateCameraZoom() {
	if rl.IsKeyPressed(rl.KeyMinus) || rl.GetMouseWheelMoveV().Y < 0 {
		options.Zoom = max(4, options.Zoom-1)
	}
	if rl.IsKeyPressed(rl.KeyEqual) || rl.GetMouseWheelMoveV().Y > 0 {
		options.Zoom = min(12, options.Zoom+1)
	}
}

func createCamera(world *World) rl.Camera2D {
	target := world.position[world.player]

	offset := rl.Vector2{X: 0, Y: 0}
	if world.uistate.cursorHideTimer < CursorHideCooldown {
		offset = rl.GetMousePosition()
		offset = rl.Vector2Subtract(offset, util.HalfScreenSize())
		offset = rl.Vector2Scale(offset, 1/options.Zoom/4)
	}
	world.uistate.cameraOffset = rl.Vector2Lerp(world.uistate.cameraOffset, offset, 0.2)

	target = rl.Vector2Add(target, world.uistate.cameraOffset)
	camera := rl.Camera2D{
		Target: target,
		Offset: util.HalfScreenSize(),
		Zoom:   options.Zoom,
	}
	return camera
}
