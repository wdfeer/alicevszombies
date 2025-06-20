package internal

import (
	"alicevszombies/internal/util"
	"os"
	"strconv"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var stats = struct {
	TimePlayed float32
}{
	0,
}

var statAutosaveTimer float32 = 0

func updateStats(world *World) {
	stats.TimePlayed += dt

	statAutosaveTimer += dt
	if statAutosaveTimer >= 15 {
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
	timePlayed := "Time played: " + strconv.FormatFloat(float64(stats.TimePlayed), 'f', 2, 64)
	raygui.Label(util.RectangleV(origin, size), timePlayed)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontsize)
}
