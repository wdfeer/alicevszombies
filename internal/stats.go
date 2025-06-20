package internal

import (
	"alicevszombies/internal/util"
	"fmt"
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var stats = struct {
	TimePlayed    float32
	EnemiesKilled uint
	DollsSummoned uint
	HighestWave   map[Difficulty]uint
	RunCount      map[Difficulty]uint
}{
	HighestWave: make(map[Difficulty]uint),
	RunCount:    make(map[Difficulty]uint),
}

var statAutosaveTimer float32 = 0

func updateStats(world *World) {
	stats.TimePlayed += dt

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

func renderStats(origin rl.Vector2) {
	size := rl.Vector2{X: 480, Y: 120}
	spacing := float32(40)
	panelSize := rl.Vector2{X: size.X, Y: size.Y*4 + spacing*5}

	oldFontsize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 32)

	raygui.Panel(util.RectangleV(origin, panelSize), "")

	origin.Y += spacing
	timePlayed := "Time played: "
	if stats.TimePlayed > 60 {
		timePlayed += fmt.Sprint(int(stats.TimePlayed)/60) + "m"
		timePlayed += fmt.Sprint(int(stats.TimePlayed)%60) + "s"
	} else {
		timePlayed += fmt.Sprint(int(stats.TimePlayed)) + "s"
	}
	raygui.Label(util.RectangleV(origin, size), timePlayed)

	origin.Y += spacing
	dollsSummoned := "Dolls summoned: " + fmt.Sprint(stats.DollsSummoned)
	raygui.Label(util.RectangleV(origin, size), dollsSummoned)

	origin.Y += spacing
	killCounter := "Enemies killed: " + fmt.Sprint(stats.EnemiesKilled)
	raygui.Label(util.RectangleV(origin, size), killCounter)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
