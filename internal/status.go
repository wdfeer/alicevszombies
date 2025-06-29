package internal

type Status struct {
	poison float32
	// TODO: implement slow/weak for blue stuff
}

func updateStatus(world *World) {
	for id, status := range world.status {
		// TODO: poison deals damage and reduces duration
	}
}
