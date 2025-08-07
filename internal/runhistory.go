package internal

import (
	"alicevszombies/internal/colors"
	"alicevszombies/internal/util"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var runHistory = struct {
	Entries []RunEntry
	page    int32
}{
	Entries: make([]RunEntry, 0),
}

type RunEntry = struct {
	Time         time.Time
	Difficulty   Difficulty
	WaveReached  uint16
	DollCount    uint16
	UpgradeCount uint16
	Playtime     float32
}

const runHistoryPath = "user/runhistory.bin"

func loadRunHistory() {
	data, err := os.ReadFile(runHistoryPath)
	if err == nil {
		if err = util.Deserialize(data, &runHistory); err == nil {
			println(fmt.Sprintf("INFO: Loaded run history with %d entries successfully!", len(runHistory.Entries)))
			return
		} else {
			println("ERROR: Failed deserializing run history!")
		}
	} else {
		println("ERROR: Failed reading run history file!")
	}

	println("WARNING: Creating default run history file...")

	go saveRunHistory()
}

func saveRunHistory() {
	bytes, err := util.Serialize(&runHistory)
	if err != nil {
		println("ERROR: Failed serializing run history!")
		return
	}

	if _, err = os.Stat("user"); err != nil {
		err = os.Mkdir("user", 0755)
		if err != nil {
			println("ERROR: Failed creating \"user\" directory!")
			return
		}
	}

	err = os.WriteFile(runHistoryPath, bytes, 0644)
	if err != nil {
		println("ERROR: Failed writing run history file!")
		return
	}
	println(fmt.Sprintf("INFO: Run history saved with %d entries", len(runHistory.Entries)))
}

func SaveRun(world *World) {
	if world.enemySpawner.wave > 1 {
		runHistory.Entries = append(runHistory.Entries, RunEntry{
			Time:         time.Now(),
			Difficulty:   world.difficulty,
			WaveReached:  uint16(world.enemySpawner.wave),
			DollCount:    uint16(len(world.doll)),
			UpgradeCount: uint16(world.playerData.upgradeCount()),
			Playtime:     world.playtime,
		})

		go saveRunHistory()
	}
}

func renderRunHistory(origin rl.Vector2) {
	size := rl.Vector2{X: 720 * uiScale, Y: 120 * uiScale}
	margin := float32(20) * uiScale

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	oldLineSpacing := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING)
	oldTextAlign := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, textSize40/2)

	var startIndex int
	if len(runHistory.Entries) > 3 { // Show page switcher
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize64)
		textWidth := float32(rl.MeasureText("Page", int32(textSize64)))
		width := size.X - textWidth
		raygui.SetStyle(raygui.SPINNER, raygui.ARROWS_SIZE, int64(width)/7)
		raygui.Spinner(rl.Rectangle{X: origin.X, Y: origin.Y, Width: width, Height: size.Y + margin}, "Page", &runHistory.page, 1, max(1, len(runHistory.Entries)/3+1), false)
		startIndex = int(runHistory.page)*3 - 3
		origin.Y += size.Y + margin*2
	} else {
		startIndex = max(len(runHistory.Entries)-3, 0)
	}

	shownRuns := runHistory.Entries[startIndex:min(startIndex+3, len(runHistory.Entries))]
	for i, e := range shownRuns {
		// TODO: revert y positions to put recent runs on top
		rectY := origin.Y + float32(i)*(size.Y+margin*3) + margin
		rect := rl.Rectangle{X: origin.X + margin, Y: rectY, Width: size.X, Height: size.Y}

		raygui.Panel(rl.Rectangle{X: origin.X, Y: rectY - margin, Width: size.X + margin*2, Height: size.Y + margin*2}, "")

		// Difficulty
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
		{
			oldColor := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_NORMAL)
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
			var diffStr string
			var diffColor rl.Color
			switch e.Difficulty {
			case EASY:
				diffStr = "Easy"
				diffColor = rl.White
			case NORMAL:
				diffStr = "Normal"
				diffColor = colors.Yellow
			case HARD:
				diffStr = "Hard"
				diffColor = colors.Red
			case LUNATIC:
				diffStr = "Lunatic"
				diffColor = colors.Purple
			}
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_NORMAL, int64(rl.ColorToInt(diffColor)))
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
			raygui.Label(rect, diffStr)
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_NORMAL, oldColor)
		}

		// Wave Reached
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_RIGHT))
		waveStr := fmt.Sprintf("%d Waves", e.WaveReached)
		raygui.Label(rect, waveStr)

		// How long ago Played
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize24)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_CENTER))
		timeAgoStr := util.TimeAgo(e.Time)
		rect.Height /= 4
		raygui.Label(rect, timeAgoStr)
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, oldLineSpacing)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, oldTextAlign)
}
