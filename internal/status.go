package internal

import "alicevszombies/internal/util"

type Status struct {
	poison float32
	slow   float32
}

func updateStatus(world *World) {
	for id, status := range world.status {
		if status.poison > 0 {
			if util.ModF(status.poison, 1) < dt {
				damage(world, id, 1)
			}
		}

		world.status[id] = Status{
			poison: status.poison - dt,
			slow:   status.slow - dt,
		}
	}
}

func applyPoison(world *World, id Entity, duration float32) {
	poison := world.status[id].poison
	if duration > poison {
		world.status[id] = Status{
			poison: duration,
			slow:   world.status[id].slow,
		}
	}
}

func applySlow(world *World, id Entity, duration float32) {
	slow := world.status[id].slow
	if duration > slow {
		world.status[id] = Status{
			poison: world.status[id].poison,
			slow:   duration,
		}
	}
}
