package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Achievements = map[uint8]float32

type AchievementType struct {
	id          uint8
	name        string
	description string
}

var (
	Wave100Reached = AchievementType{
		id:          0,
		name:        "Overachiever",
		description: "Reach wave 100",
	}
)

func updateAchievements(world *World) {
	stats.Achievements[Wave100Reached.id] = float32(stats.HighestWave[UNDEFINED]) / 100
}

func renderAchievements(origin rl.Vector2) {
	size := rl.Vector2{X: 480, Y: 120}
	spacing := float32(40)
	count := len(stats.Achievements)
	panelSize := rl.Vector2{X: size.X + spacing*2, Y: size.Y*float32(count) + spacing*(float32(count)+1)}

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)

	raygui.Panel(util.RectangleV(origin, panelSize), "")

	x := origin.X + spacing
	for id, _ := range stats.Achievements {
		// TODO: get name/description by id for displaying it
		name := "TODO"
		description := "TODO"

		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)
		raygui.Label(rl.Rectangle{X: x, Y: origin.Y + (float32(id)+1)*(size.Y+spacing), Width: size.X, Height: size.Y / 2}, name)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 24)
		raygui.Label(rl.Rectangle{X: x, Y: origin.Y + (float32(id)+1)*(size.Y+spacing) + size.Y/2, Width: size.X, Height: size.Y / 2}, description)
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
