package internal

import "alicevszombies/internal/util"

type Status struct {
	poison float32
	// TODO: implement slow/weak for blue stuff
}

func updateStatus(world *World) {
	for id, status := range world.status {
		if status.poison > 0 {
			if util.ModF(status.poison, 1) < dt {
				damage(world, id, 1)
			}
			world.status[id] = Status{
				poison: status.poison - dt,
			}
		}
	}
}

func applyPoison(world *World, id Entity, duration float32) {
	poison := world.status[id].poison
	if duration > poison {
		world.status[id] = Status{
			poison: duration,
		}
	}
}
