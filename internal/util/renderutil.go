package util

import rl "github.com/gen2brain/raylib-go/raylib"

func GetHalfScreen() rl.Vector2 {
	return rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
}

func DrawTextureCentered(texture rl.Texture2D, center rl.Vector2) {
	pos := rl.Vector2Subtract(center, rl.Vector2{X: float32(texture.Width) / 2, Y: float32(texture.Height) / 2})
	rl.DrawTextureV(texture, pos, rl.White)
}
