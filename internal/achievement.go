package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Achievements = map[*AchievementType]float32

type AchievementType struct {
	name        string
	description string
}

var (
	Wave100Reached = AchievementType{
		name:        "Overachiever",
		description: "Reach wave 100",
	}
)

func updateAchievements(world *World) {
	stats.Achievements[&Wave100Reached] = float32(stats.HighestWave[UNDEFINED]) / 100
}

func renderAchievements(origin rl.Vector2) {
	size := rl.Vector2{X: 480, Y: 120}
	spacing := float32(40)
	panelSize := rl.Vector2{X: size.X, Y: size.Y*4 + spacing*5}

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)

	raygui.Panel(util.RectangleV(origin, panelSize), "")

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)

	// TODO: display achievements

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
