package internal

import (
	"alicevszombies/internal/util"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var runHistory = struct {
	Entries map[uint16]RunEntry
}{}

type RunEntry = struct {
	Difficulty   Difficulty
	WaveReached  uint16
	DollCount    uint16
	UpgradeCount uint16
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

	go saveHistory()
}

func saveRunHistory() {
	bytes, err := util.Serialize(&runHistory)
	if err != nil {
		println("ERROR: Failed serializing run history!")
		return
	}

	if_, err = os.Stat("user"); err != nil {
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

func renderRunHistory(origin rl.Vector2) {
	size := rl.Vector2{X: 720 * uiScale, Y: 480 * uiScale}
	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	oldLineSpacing := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, textSize40/2)

	raygui.Panel(util.RectangleV(origin, size), "TODO")

	// TODO: implement run rendering

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_LINE_SPACING, oldLineSpacing)
}
