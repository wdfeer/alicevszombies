package internal

import (
	"alicevszombies/internal/util"
	"os"
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
