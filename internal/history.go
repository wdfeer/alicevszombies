package internal

import (
	"alicevszombies/internal/util"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var history = struct {
	tabSelected          uint8
	TimePlayed           map[Difficulty]float32
	EnemiesKilledTotal   map[Difficulty]uint
	DollsSummoned        map[Difficulty]uint
	HighestWave          map[Difficulty]uint
	RunCount             map[Difficulty]uint
	Achievements         Achievements
	UpgradesUsed         map[string]uint
	EnemiesKilledPerType map[string]uint
}{
	TimePlayed:           make(map[Difficulty]float32),
	EnemiesKilledTotal:   make(map[Difficulty]uint),
	DollsSummoned:        make(map[Difficulty]uint),
	HighestWave:          make(map[Difficulty]uint),
	RunCount:             make(map[Difficulty]uint),
	Achievements:         Achievements{},
	UpgradesUsed:         make(map[string]uint),
	EnemiesKilledPerType: make(map[string]uint),
}

var historyAutosaveTimer float32 = 0

func updateHistory(world *World) {
	history.TimePlayed[world.difficulty] += dt

	if world.enemySpawner.wave > uint32(history.HighestWave[world.difficulty]) {
		history.HighestWave[world.difficulty] = uint(world.enemySpawner.wave)
	}

	historyAutosaveTimer += dt
	if historyAutosaveTimer >= 15 {
		historyAutosaveTimer = 0
		go saveHistory()
	}

	updateAchievements(world)
}

const historyPath = "user/history.bin"

func loadHistory() {
	data, err := os.ReadFile(historyPath)
	if err == nil {
		if err = util.Deserialize(data, &history); err == nil {
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

func saveHistory() {
	bytes, err := util.Serialize(&history)
	if err != nil {
		println("ERROR: Failed serializing history!")
		return
	}

	if _, err = os.Stat("user"); err != nil {
		err = os.Mkdir("user", 0755)
		if err != nil {
			println("ERROR: Failed creating \"user\" directory!")
			return
		}
	}

	err = os.WriteFile(historyPath, bytes, 0644)
	if err != nil {
		println("ERROR: Failed writing history file!")
		return
	}
	println("INFO: History saved!")
}

func renderHistory(origin rl.Vector2) {

	{ // Tabs
		o := origin
		oldTextSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldTextSize/2)
		buttonWidth := 240. * uiScale
		buttonSpacing := 60 * uiScale
		buttonHeight := 30 * uiScale
		if raygui.Toggle(rl.Rectangle{X: o.X, Y: o.Y, Width: buttonWidth, Height: buttonHeight}, "Stats", history.tabSelected == 0) {
			history.tabSelected = 0
		}
		o.X += buttonWidth + buttonSpacing
		if raygui.Toggle(rl.Rectangle{X: o.X, Y: o.Y, Width: buttonWidth, Height: buttonHeight}, "Achievements", history.tabSelected == 1) {
			history.tabSelected = 1
		}
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldTextSize)
		origin.Y += buttonHeight + buttonSpacing
	}

	switch history.tabSelected {
	case 0:
		renderStats(origin)
	case 1:
		renderAchievements(origin)
	}
}
