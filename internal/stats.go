package internal

import (
	"alicevszombies/internal/util"
	"fmt"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var stats = struct {
	TimePlayed    map[Difficulty]float32
	EnemiesKilled map[Difficulty]uint
	DollsSummoned map[Difficulty]uint
	HighestWave   map[Difficulty]uint
	RunCount      map[Difficulty]uint
}{
	TimePlayed:    make(map[Difficulty]float32),
	EnemiesKilled: make(map[Difficulty]uint),
	DollsSummoned: make(map[Difficulty]uint),
	HighestWave:   make(map[Difficulty]uint),
	RunCount:      make(map[Difficulty]uint),
}

var statAutosaveTimer float32 = 0

func updateStats(world *World) {
	stats.TimePlayed[world.difficulty] += dt

	if world.enemySpawner.wave > uint32(stats.HighestWave[world.difficulty]) {
		stats.HighestWave[world.difficulty] = uint(world.enemySpawner.wave)
	}

	statAutosaveTimer += dt
	if statAutosaveTimer >= 15 {
		statAutosaveTimer = 0
		go saveStats()
	}
}

func loadStats() {
	data, err := os.ReadFile("user/stats.bin")
	if err == nil {
		if err = util.Deserialize(data, &stats); err == nil {
			println("INFO: Loaded stats successfully!")
			return
		} else {
			println("ERROR: Failed deserializing stats!")
		}
	} else {
		println("ERROR: Failed reading stats file!")
	}

	println("WARNING: Creating default stats file...")

	go saveStats()
}

func saveStats() {
	bytes, err := util.Serialize(&stats)
	if err != nil {
		println("ERROR: Failed serializing stats!")
		return
	}

	if _, err = os.Stat("user"); err != nil {
		err = os.Mkdir("user", 0755)
		if err != nil {
			println("ERROR: Failed creating \"user\" directory!")
			return
		}
	}

	err = os.WriteFile("user/stats.bin", bytes, 0644)
	if err != nil {
		println("ERROR: Failed writing stats file!")
		return
	}
	println("INFO: Stats saved!")
}

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
		for _, v := range stats.TimePlayed {
			timePlayed += v
		}
	} else {
		timePlayed = stats.TimePlayed[statSelectedDifficulty]
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
		for _, v := range stats.DollsSummoned {
			dollsSummoned += v
		}
	} else {
		dollsSummoned = stats.DollsSummoned[statSelectedDifficulty]
	}
	dollsSummonedText := "Dolls summoned: " + fmt.Sprint(dollsSummoned)
	raygui.Label(util.RectangleV(origin, size), dollsSummonedText)

	origin.Y += spacing
	kills := uint(0)
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range stats.EnemiesKilled {
			kills += v
		}
	} else {
		kills = stats.EnemiesKilled[statSelectedDifficulty]
	}
	killCounter := "Enemies killed: " + fmt.Sprint(kills)
	raygui.Label(util.RectangleV(origin, size), killCounter)

	origin.Y += spacing
	var highestWave uint
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range stats.HighestWave {
			if v > highestWave {
				highestWave = v
			}
		}
	} else {
		highestWave = stats.HighestWave[statSelectedDifficulty]
	}
	highestWaveText := "Highest wave: " + fmt.Sprint(highestWave)
	raygui.Label(util.RectangleV(origin, size), highestWaveText)

	origin.Y += spacing
	var runCount uint
	if statSelectedDifficulty == UNDEFINED {
		for _, v := range stats.RunCount {
			runCount += v
		}
	} else {
		runCount = stats.RunCount[statSelectedDifficulty]
	}
	runCountText := "Run count: " + fmt.Sprint(runCount)
	raygui.Label(util.RectangleV(origin, size), runCountText)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
