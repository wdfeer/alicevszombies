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
			println("INFO: Loaded history successfully!")
			return
		} else {
			println("ERROR: Failed deserializing history!")
		}
	} else {
		println("ERROR: Failed reading history file!")
	}

	println("WARNING: Creating default history file...")

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
	println("INFO: Run history saved!")
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
	}
}

func renderRunHistory(origin rl.Vector2) {
	size := rl.Vector2{X: 720 * uiScale, Y: 120 * uiScale}
	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	oldLineSpacing := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING)
	oldTextAlign := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, textSize40/2)

	margin := float32(20) * uiScale
	startIndex := max(len(runHistory.Entries)-3, 0)
	runsShown := len(runHistory.Entries) - startIndex
	for i := range min(len(runHistory.Entries), 3) {
		index := startIndex + i
		rectY := origin.Y + float32((3-i)-(4-runsShown))*(size.Y+margin*3) + margin // Surely it doesn't need to be so complicated...
		rect := rl.Rectangle{X: origin.X + margin, Y: rectY, Width: size.X, Height: size.Y}

		{ // Background panel
			panelRect := rect
			panelRect.X -= margin
			panelRect.Y -= margin
			panelRect.Width += margin * 2
			panelRect.Height += margin * 2
			raygui.Panel(panelRect, "")
		}

		e := &runHistory.Entries[index]

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
