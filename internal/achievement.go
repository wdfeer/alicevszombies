package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Achievements = [5]float32

type AchievementType struct {
	id                uint8
	name              string
	description       string
	visualMaxProgress float32
}

var (
	AllUpgradesObtained = AchievementType{
		id:          0,
		name:        "Arcane Arsenal",
		description: "Obtain each upgrade once",
	}
	AllEnemiesKilled = AchievementType{
		id:          1,
		name:        "Apparition Expert",
		description: "Defeat every enemy once",
	}
	Wave30OneDoll = AchievementType{
		id:                2,
		name:              "Solo Marionette",
		description:       "Reach Wave 30 while\nhaving only one Doll",
		visualMaxProgress: 1,
	}
	Wave30Lunatic = AchievementType{
		id:                3,
		name:              "Crimson Thread",
		description:       "Reach Wave 30 on Lunatic",
		visualMaxProgress: 30,
	}
	Wave50Reached = AchievementType{
		id:                4,
		name:              "Zombie Slayer",
		description:       "Reach Wave 50",
		visualMaxProgress: 50,
	}
)

func initAchievements() {
	AllUpgradesObtained.visualMaxProgress = float32(len(upgrades) + len(uniqueUpgrades))
	AllEnemiesKilled.visualMaxProgress = float32(len(allEnemyTypes))
}

func updateAchievements(world *World) {
	oldAchievements := history.Achievements

	if world.enemySpawner.wave >= 30 && len(world.doll) == 1 {
		history.Achievements[Wave30OneDoll.id] = 1
	}

	history.Achievements[Wave30Lunatic.id] = float32(history.HighestWave[LUNATIC]) / 30

	var highestWave uint
	for _, v := range history.HighestWave {
		if v > highestWave {
			highestWave = v
		}
	}
	history.Achievements[Wave50Reached.id] = float32(highestWave) / 50

	history.Achievements[AllUpgradesObtained.id] = float32(len(history.UpgradesUsed)) / AllUpgradesObtained.visualMaxProgress

	history.Achievements[AllEnemiesKilled.id] = float32(len(history.EnemiesKilledPerType)) / AllEnemiesKilled.visualMaxProgress

	if history.Achievements != oldAchievements {
		for i, v := range history.Achievements {
			if v >= 1 && oldAchievements[i] < 1 {
				showAchievementNotification(world, uint8(i))
			}
		}
	}
}

var achievementsByID = map[uint8]*AchievementType{
	Wave30OneDoll.id:       &Wave30OneDoll,
	Wave30Lunatic.id:       &Wave30Lunatic,
	Wave50Reached.id:       &Wave50Reached,
	AllUpgradesObtained.id: &AllUpgradesObtained,
	AllEnemiesKilled.id:    &AllEnemiesKilled,
}

func renderAchievements(origin rl.Vector2) {
	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	oldLineSpacing := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, textSize40/2)

	size := rl.Vector2{X: 760 * uiScale, Y: 120 * uiScale}
	margin := float32(20) * uiScale

	for id := range achievementsByID {
		origin := rl.Vector2{X: origin.X, Y: origin.Y + float32(id)*(size.Y+margin*3)}
		renderAchievement(origin, size, margin, id)
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, oldLineSpacing)
}

func renderAchievement(origin rl.Vector2, size rl.Vector2, margin float32, achievementID uint8) {
	ach := achievementsByID[achievementID]

	progress := history.Achievements[achievementID]
	panelRect := util.RectangleV(origin, size)

	if progress >= 1 {
		raygui.SetState(raygui.STATE_FOCUSED)
	}

	raygui.Panel(panelRect, "")

	rect := rl.Rectangle{X: panelRect.X + margin, Y: panelRect.Y + margin, Width: panelRect.Width - margin*2, Height: panelRect.Height - margin*2}

	// Title
	rect.Height /= 4
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
	raygui.Label(rect, ach.name)

	// Description
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize24)
	rect.Y += size.Y / 4
	rect.Height = size.Y / 2
	raygui.Label(rect, ach.description)

	// Progress
	rect.X = panelRect.X
	rect.Width = panelRect.Width
	rect.Y += size.Y/2 + margin/2
	rect.Height = size.Y / 4
	raygui.ProgressBar(rect, "", "", progress, 0, 1)
	raygui.Label(rect, fmt.Sprint(progress*ach.visualMaxProgress)+"/"+fmt.Sprint(ach.visualMaxProgress))

	if progress >= 1 {
		raygui.SetState(raygui.STATE_NORMAL)
	}
}

type AchievementNotification = struct {
	id       uint8
	timeLeft float32
}

func showAchievementNotification(world *World, achievementID uint8) {
	world.uistate.achievementNotification = AchievementNotification{
		id:       achievementID,
		timeLeft: 1.5,
	}
}

func renderAchievementNotification(world *World) {
	if world.uistate.achievementNotification.timeLeft <= 0 {
		return
	}

	size := rl.Vector2{X: 560 * uiScale, Y: 120 * uiScale}
	margin := float32(16) * uiScale
	renderAchievement(rl.Vector2Subtract(util.ScreenSize(), rl.Vector2Add(size, rl.Vector2{X: margin * 2, Y: margin * 2})), size, margin, world.uistate.achievementNotification.id)
	world.uistate.achievementNotification.timeLeft -= dt
}
