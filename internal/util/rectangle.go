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

// Same as rl.CheckCollisionRecs. Used to avoid C calls for performance.
func CheckCollisionRecs(rec1 rl.Rectangle, rec2 rl.Rectangle) bool {
	return rec1.X+rec1.Width >= rec2.X &&
		rec1.X <= rec2.X+rec2.Width &&
		rec1.Y+rec1.Height >= rec2.Y &&
		rec1.Y <= rec2.Y+rec2.Height
}
