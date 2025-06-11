package internal

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HP struct {
	val              float32
	attackerCooldown map[Entity]float32
}

func newHP(amount float32) HP {
	return HP{
		val:              amount,
		attackerCooldown: make(map[Entity]float32),
	}
}

func updateHP(world *World) {
	for id, hp := range world.hp {
		newMap := hp.attackerCooldown

		for id := range newMap {
			newMap[id] -= rl.GetFrameTime()
		}

		world.hp[id] = HP{
			val:              hp.val,
			attackerCooldown: newMap,
		}

	}
}

func damage(world *World, id Entity, amount float32) {
	hp := world.hp[id]
	hp.val -= amount

	newCombatText(world, world.position[id], fmt.Sprint(amount))

	if hp.val <= 0 {
		println("Entity with id", id, "killed!")
		world.deleteEntity(id)
	} else {
		world.hp[id] = hp
	}
}

func damageWithCooldown(world *World, id Entity, amount float32, attacker Entity) {
	hp := world.hp[id]
	if cooldown, exists := hp.attackerCooldown[attacker]; !exists || cooldown <= 0 {
		hp.attackerCooldown[attacker] = 0.5
		world.hp[id] = hp
		damage(world, id, amount)
	}
}
