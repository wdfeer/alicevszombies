package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateMainMenu(world *World) {
	switch rl.GetKeyPressed() {
	case rl.KeyOne:
		world.difficulty = EASY
	case rl.KeyTwo:
		world.difficulty = NORMAL
	case rl.KeyThree:
		world.difficulty = HARD
	case rl.KeyFour:
		world.difficulty = LUNATIC
	default:
		return
	}
	world.paused = false
	world.uistate.isMainMenu = false
}

func renderMainMenu() {
	screenSize := util.ScreenSize()
	rl.DrawRectangleV(rl.Vector2Zero(), screenSize, rl.ColorAlpha(rl.Black, 0.9))

	util.DrawTextCenteredSpaced("alicevszombies", 80, rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 120}, 16)

	{ // Left Side
		height := int32(screenSize.Y)
		rl.DrawText("Fullscreen = F", 160, height-160, 64, rl.White)
		rl.DrawText("Exit = DEL", 160, height-96, 64, rl.White)
	}

	{ // Right Side
		spacing := float32(8)
		pos := util.AlignTextRight("Easy = 1", 64, rl.Vector2{X: screenSize.X - 160, Y: screenSize.Y - 64*4}, spacing)
		rl.DrawTextEx(rl.GetFontDefault(), "Easy = 1", pos, 64, spacing, rl.White)
		pos = util.AlignTextRight("Normal = 2", 64, rl.Vector2{X: screenSize.X - 160, Y: screenSize.Y - 64*3}, spacing)
		rl.DrawTextEx(rl.GetFontDefault(), "Normal = 2", pos, 64, spacing, rl.White)
		pos = util.AlignTextRight("Hard = 3", 64, rl.Vector2{X: screenSize.X - 160, Y: screenSize.Y - 64*2}, spacing)
		rl.DrawTextEx(rl.GetFontDefault(), "Hard = 3", pos, 64, spacing, rl.White)
		pos = util.AlignTextRight("Lunatic = 4", 64, rl.Vector2{X: screenSize.X - 160, Y: screenSize.Y - 64}, spacing)
		rl.DrawTextEx(rl.GetFontDefault(), "Lunatic = 4", pos, 64, spacing, rl.White)
	}
}
