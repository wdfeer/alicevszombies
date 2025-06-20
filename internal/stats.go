package internal

var stats = struct {
	TimePlayed float32
}{
	0,
}

func updateStats(world *World) {
	stats.TimePlayed += dt
}

func loadStats() {
	// TODO: deserialize
}

func saveStats() {
	// TODO: serialize
}
