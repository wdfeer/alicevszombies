package internal

import (
	"alicevszombies/internal/util"
	"math/rand"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	selected     uint8
	dollPosition [8]rl.Vector2
	dollVelocity [8]rl.Vector2
}

func renderMainMenu(world *World) {
	screenSize := util.ScreenSize()
	rl.DrawRectangleV(rl.Vector2Zero(), screenSize, rl.ColorAlpha(rl.Black, 0.9))

	buttonWidth := float32(480) * uiScale
	buttonHeight := float32(120) * uiScale
	buttonSpacing := float32(40) * uiScale
	startY := buttonHeight + buttonSpacing

	x := screenSize.X / 20
	y := startY

	mainMenu := &world.uistate.mainMenu

	if raygui.LabelButton(rl.Rectangle{X: screenSize.X - buttonWidth*1.5, Y: y, Width: buttonWidth * 1.5, Height: buttonHeight}, "alicevszombies") {
		for i, p := range mainMenu.dollPosition {
			if p == rl.Vector2Zero() {
				mainMenu.dollPosition[i] = rl.Vector2{X: screenSize.X * rand.Float32(), Y: -20}
				mainMenu.dollVelocity[i] = rl.Vector2{X: (rand.Float32() - 0.5) * screenSize.X, Y: 0}
				break
			}
		}
	}
	for i, p := range mainMenu.dollPosition {
		if p != rl.Vector2Zero() {
			texture := "doll_sword"
			if i%5 == 4 {
				texture = "doll_magician"
			}
			rl.DrawTextureEx(assets.textures[texture], p, 0, 8, rl.White)
			mainMenu.dollVelocity[i] = rl.Vector2Add(mainMenu.dollVelocity[i], rl.Vector2{X: 0, Y: screenSize.Y * dt * 4})
			mainMenu.dollPosition[i] = rl.Vector2Add(p, rl.Vector2Scale(mainMenu.dollVelocity[i], dt))
			if mainMenu.dollPosition[i].Y > screenSize.Y+100 {
				mainMenu.dollPosition[i] = rl.Vector2Zero()
			}
		}
	}

	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Start", mainMenu.selected == 1) {
		mainMenu.selected = 1
	} else if mainMenu.selected == 1 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "History", mainMenu.selected == 2) {
		mainMenu.selected = 2
	} else if mainMenu.selected == 2 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Toggle(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Options", mainMenu.selected == 3) {
		mainMenu.selected = 3
	} else if mainMenu.selected == 3 {
		mainMenu.selected = 0
	}

	y += buttonHeight + buttonSpacing
	if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Exit") {
		rl.CloseWindow()
	}

	x += buttonWidth * 1.1
	y = startY - buttonHeight/2
	switch mainMenu.selected {
	case 1:
		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Easy") {
			startGame(world, EASY)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Normal") {
			startGame(world, NORMAL)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Hard") {
			startGame(world, HARD)
		}
		y += buttonHeight + buttonSpacing

		if raygui.Button(rl.Rectangle{X: x, Y: y, Width: buttonWidth, Height: buttonHeight}, "Lunatic") {
			startGame(world, LUNATIC)
		}
	case 2:
		renderHistory(rl.Vector2{X: x, Y: y})
	case 3:
		renderOptions(rl.Vector2{X: x, Y: y})
	}
}

func startGame(world *World, difficulty Difficulty) {
	selectDifficulty(world, difficulty)
	world.paused = false
	world.uistate.isMainMenu = false
}
