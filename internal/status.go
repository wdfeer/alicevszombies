package internal

import (
	"alicevszombies/internal/util"
)

type Status = [3]float32
type StatusType = uint

const (
	Poison StatusType = iota
	Slow
	Bleed
)

func updateStatus(world *World) {
	for id, status := range world.status {
		if status[Poison] > 0 && util.ModF(status[Poison], 1) < dt {
			damage(world, id, 1)
		}
		if status[Bleed] > 0 && util.ModF(status[Bleed], 2) < dt {
			damage(world, id, 1)
		}

		newStatus := status
		for i := range newStatus {
			newStatus[i] -= dt
		}
		world.status[id] = newStatus
	}
}

func applyStatus(world *World, id Entity, typ StatusType, duration float32) {
	newStatus := world.status[id]
	newStatus[typ] = max(newStatus[typ], duration)
	world.status[id] = newStatus
}
