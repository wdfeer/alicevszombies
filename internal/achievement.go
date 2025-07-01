package internal

import rl "github.com/gen2brain/raylib-go/raylib"

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
	// TODO: render achievement panel
}
