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
		var damageOverTime float32
		if world.enemySpawner.wave > 30 {
			damageOverTime = 1 + float32((world.enemySpawner.wave-30)/20)
		} else {
			damageOverTime = 1
		}

		if status[Poison] > 0 && util.ModF(status[Poison], 1) < dt {
			damage(world, id, damageOverTime)
		}
		if status[Bleed] > 0 && util.ModF(status[Bleed], 2) < dt {
			damage(world, id, damageOverTime)
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
