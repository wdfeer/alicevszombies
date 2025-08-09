package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var statSelectedDifficulty Difficulty = UNDEFINED

func renderStats(origin rl.Vector2) { // TODO: refactor this monstrosity of a function
	size := rl.Vector2{X: 760 * uiScale, Y: 250 * uiScale}
	diffSelectorHeight := 120 * uiScale
	spacing := float32(40) * uiScale

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	oldLineSpacing := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING)

	panelPos := rl.Vector2{X: origin.X, Y: origin.Y + diffSelectorHeight + spacing}
	panelSize := rl.Vector2{X: size.X, Y: size.Y}
	raygui.Panel(util.RectangleV(panelPos, panelSize), "")

	diffText := "Overall\nEasy\nNormal\nHard\nLunatic"

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize64)
	raygui.SetStyle(raygui.COMBOBOX, raygui.COMBO_BUTTON_WIDTH, int64(size.X)/4)
	statSelectedDifficulty = Difficulty(raygui.ComboBox(util.RectangleV(origin, rl.Vector2{X: size.X, Y: diffSelectorHeight}), diffText, int32(statSelectedDifficulty)))

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, textSize40)

	str := ""

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
	str += timePlayedText + "\n"

	dollsSummoned := uint(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.DollsSummoned {
			dollsSummoned += v
		}
	} else {
		dollsSummoned = history.DollsSummoned[statSelectedDifficulty]
	}
	dollsSummonedText := "Dolls summoned: " + fmt.Sprint(dollsSummoned)
	str += dollsSummonedText + "\n"

	kills := uint(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.EnemiesKilledTotal {
			kills += v
		}
	} else {
		kills = history.EnemiesKilledTotal[statSelectedDifficulty]
	}
	killCounter := "Enemies killed: " + fmt.Sprint(kills)
	str += killCounter + "\n"

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
	str += highestWaveText + "\n"

	var runCount uint
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range history.RunCount {
			runCount += v
		}
	} else {
		runCount = history.RunCount[statSelectedDifficulty]
	}
	runCountText := "Run count: " + fmt.Sprint(runCount)
	str += runCountText + "\n"

	panelPos.Y += float32(textSize40) * 1.75 // The text is too high otherwise. IDK why
	raygui.Label(util.RectangleV(panelPos, panelSize), str)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldLineSpacing)
}
