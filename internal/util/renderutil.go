package util

import rl "github.com/gen2brain/raylib-go/raylib"

func GetHalfScreen() rl.Vector2 {
	return rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
}

func CenterSomething(width float32, height float32, desiredCenter rl.Vector2) rl.Vector2 {
	return rl.Vector2Subtract(desiredCenter, rl.Vector2{X: width / 2, Y: height / 2})
}

func CenterTexture(texture *rl.Texture2D, desiredCenter rl.Vector2) rl.Vector2 {
	return rl.Vector2Subtract(desiredCenter, rl.Vector2{X: float32(texture.Width) / 2, Y: float32(texture.Height) / 2})
}

func DrawTextureCentered(texture rl.Texture2D, center rl.Vector2) {
	rl.DrawTextureV(texture, CenterTexture(&texture, center), rl.White)
}

func DrawTextureCenteredScaled(texture rl.Texture2D, center rl.Vector2, scale float32) {
	pos := rl.Vector2Subtract(center, rl.Vector2{X: float32(texture.Width) * scale / 2, Y: float32(texture.Height) * scale / 2})
	rl.DrawTextureEx(texture, pos, 0, scale, rl.White)
}

func CenterText(text string, fontSize float32, desiredCenter rl.Vector2) rl.Vector2 {
	size := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, 0)
	return rl.Vector2Subtract(desiredCenter, rl.Vector2Scale(size, 0.5))
}

func CenterTextSpaced(text string, fontSize float32, desiredCenter rl.Vector2, spacing float32) rl.Vector2 {
	size := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, spacing)
	return rl.Vector2Subtract(desiredCenter, rl.Vector2Scale(size, 0.5))
}

func DrawTextCentered(text string, fontSize float32, center rl.Vector2) {
	pos := CenterText(text, fontSize, center)
	rl.DrawTextEx(rl.GetFontDefault(), text, pos, fontSize, 0, rl.White)
}

func DrawTextCenteredSpaced(text string, fontSize float32, center rl.Vector2, spacing float32) {
	pos := CenterTextSpaced(text, fontSize, center, spacing)
	rl.DrawTextEx(rl.GetFontDefault(), text, pos, fontSize, spacing, rl.White)
}
