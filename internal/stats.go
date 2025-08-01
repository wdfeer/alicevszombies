package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var statSelectedDifficulty Difficulty = UNDEFINED

func renderStats(origin rl.Vector2) { // TODO: refactor this monstrosity of a function
	size := rl.Vector2{X: 480, Y: 120}
	spacing := float32(40)
	panelSize := rl.Vector2{X: size.X, Y: size.Y*4 + spacing*5}

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)

	raygui.Panel(util.RectangleV(origin, panelSize), "")

	origin.Y += spacing
	diffText := "Overall\nEasy\nNormal\nHard\nLunatic"

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 48)
	raygui.SetStyle(raygui.COMBOBOX, raygui.COMBO_BUTTON_WIDTH, int64(size.X)/4)
	statSelectedDifficulty = Difficulty(raygui.ComboBox(util.RectangleV(origin, size), diffText, int32(statSelectedDifficulty)))

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)

	origin.Y += spacing + size.Y
	timePlayed := float32(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.TimePlayed {
			timePlayed += v
		}
	} else {
		timePlayed = history.TimePlayed[statSelectedDifficulty]
	}
	timePlayedText := "Time played: "
	if timePlayed > 60 {
		timePlayedText += fmt.Sprint(int(timePlayed)/60) + "m"
		timePlayedText += fmt.Sprint(int(timePlayed)%60) + "s"
	} else {
		timePlayedText += fmt.Sprint(int(timePlayed)) + "s"
	}
	raygui.Label(util.RectangleV(origin, size), timePlayedText)

	origin.Y += spacing
	dollsSummoned := uint(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.DollsSummoned {
			dollsSummoned += v
		}
	} else {
		dollsSummoned = history.DollsSummoned[statSelectedDifficulty]
	}
	dollsSummonedText := "Dolls summoned: " + fmt.Sprint(dollsSummoned)
	raygui.Label(util.RectangleV(origin, size), dollsSummonedText)

	origin.Y += spacing
	kills := uint(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.EnemiesKilledTotal {
			kills += v
		}
	} else {
		kills = history.EnemiesKilledTotal[statSelectedDifficulty]
	}
	killCounter := "Enemies killed: " + fmt.Sprint(kills)
	raygui.Label(util.RectangleV(origin, size), killCounter)

	origin.Y += spacing
	var highestWave uint
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.HighestWave {
			if v > highestWave {
				highestWave = v
			}
		}
	} else {
		highestWave = history.HighestWave[statSelectedDifficulty]
	}
	highestWaveText := "Highest wave: " + fmt.Sprint(highestWave)
	raygui.Label(util.RectangleV(origin, size), highestWaveText)

	origin.Y += spacing
	var runCount uint
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.RunCount {
			runCount += v
		}
	} else {
		runCount = history.RunCount[statSelectedDifficulty]
	}
	runCountText := "Run count: " + fmt.Sprint(runCount)
	raygui.Label(util.RectangleV(origin, size), runCountText)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
