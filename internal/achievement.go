package internal

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Achievements = []float32

type AchievementType struct {
	id          uint8
	name        string
	description string
}

var (
	Wave30OneDoll = AchievementType{
		id:          0,
		name:        "Perfect Servant",
		description: "Reach Wave 30 while\nhaving only one Doll",
	}
	Wave50Lunatic = AchievementType{
		id:          1,
		name:        "Master Puppeteer",
		description: "Reach Wave 50 on Lunatic",
	}
	Wave100Reached = AchievementType{
		id:          2,
		name:        "Youkai Exterminator",
		description: "Reach Wave 100",
	}
)

func updateAchievements(world *World) {
	if world.enemySpawner.wave >= 30 && len(world.doll) == 1 {
		stats.Achievements[Wave30OneDoll.id] = 1
	}

	stats.Achievements[Wave50Lunatic.id] = float32(stats.HighestWave[LUNATIC]) / 50

	var highestWave uint
	for _, v := range stats.HighestWave {
		if v > highestWave {
			highestWave = v
		}
	}
	stats.Achievements[Wave100Reached.id] = float32(highestWave) / 50
}

var achievementsByID = map[uint8]*AchievementType{
	Wave30OneDoll.id:  &Wave30OneDoll,
	Wave50Lunatic.id:  &Wave50Lunatic,
	Wave100Reached.id: &Wave100Reached,
}

func renderAchievements(origin rl.Vector2) {
	size := rl.Vector2{X: 480, Y: 120}
	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)

	margin := float32(20)
	for id, progress := range stats.Achievements {
		rect := rl.Rectangle{X: origin.X, Y: origin.Y + float32(id)*(size.Y+margin*3), Width: size.X, Height: size.Y}

		{ // Background panel
			panelRect := rect
			panelRect.X -= margin
			panelRect.Y -= margin
			panelRect.Width += margin * 2
			panelRect.Height += margin * 2
			raygui.Panel(panelRect, "")
		}

		// Title
		rect.Height /= 4
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 36)
		raygui.Label(rect, achievementsByID[uint8(id)].name)

		// Description
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 24)
		rect.Y += size.Y / 4
		rect.Height = size.Y / 2
		raygui.Label(rect, achievementsByID[uint8(id)].description)

		// Progress
		rect.Y += size.Y / 2
		rect.Height = size.Y / 4
		raygui.ProgressBar(rect, "", "", progress, 0, 1)
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
