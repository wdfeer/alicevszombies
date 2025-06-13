package util

import rl "github.com/gen2brain/raylib-go/raylib"

func GetHalfScreen() rl.Vector2 {
	return rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
}

func DrawTextureCentered(texture rl.Texture2D, center rl.Vector2) {
	pos := rl.Vector2Subtract(center, rl.Vector2{X: float32(texture.Width) / 2, Y: float32(texture.Height) / 2})
	rl.DrawTextureV(texture, pos, rl.White)
}

func DrawTextureCenteredScaled(texture rl.Texture2D, center rl.Vector2, scale float32) {
	pos := rl.Vector2Subtract(center, rl.Vector2{X: float32(texture.Width) * scale / 2, Y: float32(texture.Height) * scale / 2})
	rl.DrawTextureEx(texture, pos, 0, scale, rl.White)
}

func CenterText(text string, fontSize float32, desiredCenter rl.Vector2) rl.Vector2 {
	size := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, 0)
	return rl.Vector2Subtract(desiredCenter, rl.Vector2Scale(size, 0.5))
}
