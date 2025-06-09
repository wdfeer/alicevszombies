package util

import rl "github.com/gen2brain/raylib-go/raylib"

func GetHalfScreen() rl.Vector2 {
	return rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
}

func DrawTextureCentered(texture rl.Texture2D, center rl.Vector2) {
	rl.DrawTexture(texture, int32(center.X)-texture.Width/2, int32(center.Y)-texture.Height/2, rl.White)
}
