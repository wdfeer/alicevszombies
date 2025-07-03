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

// Combination of CenterRectangle and CheckCollisionRecs.
func CheckCollisionCentered(pos1, size1, pos2, size2 rl.Vector2) bool {
	x1 := pos1.X - size1.X/2
	y1 := pos1.Y - size1.Y/2
	x2 := pos2.X - size2.X/2
	y2 := pos2.Y - size2.Y/2

	return x1 < x2+size2.X && x1+size1.X > x2 &&
		y1 < y2+size2.Y && y1+size1.Y > y2
}

// Combination of CenterRectangle and CheckCollisionRecs, with other rect precalculated.
func CheckCollisionCenteredVsRec(pos, size rl.Vector2, other rl.Rectangle) bool {
	x := pos.X - size.X/2
	y := pos.Y - size.Y/2

	return x < other.X+other.Width && x+size.X > other.X &&
		y < other.Y+other.Height && y+size.Y > other.Y
}
