package internal

import (
	"fmt"
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
			newMap[id] -= dt
		}

		world.hp[id] = HP{
			val:              hp.val,
			attackerCooldown: newMap,
		}

	}
}

func damage(world *World, id Entity, dmg float32) {
	hp := world.hp[id]
	hp.val -= dmg

	newCombatText(world, world.position[id], fmt.Sprint(dmg))

	if hp.val <= 0 {
		println("Entity with id", id, "killed!")
		world.deleteEntity(id)
	} else {
		world.hp[id] = hp
	}
}

func damageWithCooldown(world *World, id Entity, dmg float32, attacker Entity) {
	hp := world.hp[id]
	if cooldown, exists := hp.attackerCooldown[attacker]; !exists || cooldown <= 0 {
		hp.attackerCooldown[attacker] = 0.5
		world.hp[id] = hp
		damage(world, id, dmg)
	}
}
