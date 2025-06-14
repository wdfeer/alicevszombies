package util

import rl "github.com/gen2brain/raylib-go/raylib"

func CenterRectangle(center rl.Vector2, size rl.Vector2) rl.Rectangle {
	pos := rl.Vector2Subtract(center, rl.Vector2Scale(size, 0.5))
	return RectangleV(pos, size)
}

func RectangleV(position rl.Vector2, size rl.Vector2) rl.Rectangle {
	return rl.Rectangle{
		X:      position.X,
		Y:      position.Y,
		Width:  size.X,
		Height: size.Y,
	}
}
