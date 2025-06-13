package internal

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HP struct {
	val              float32
	attackerCooldown map[Entity]float32
	immuneTime       float32
}

func newHP(amount float32) HP {
	return HP{
		val:              amount,
		attackerCooldown: make(map[Entity]float32),
		immuneTime:       0.5,
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
			immuneTime:       hp.immuneTime,
		}
	}
}

func damage(world *World, id Entity, dmg float32) {
	hp := world.hp[id]
	hp.val -= dmg

	ctextID := newCombatText(world, world.position[id], fmt.Sprint(dmg))
	if id == world.player {
		world.combatText[ctextID] = CombatText{
			text: world.combatText[ctextID].text,
			hue:  rl.Red,
		}
	}

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
		hp.attackerCooldown[attacker] = hp.immuneTime
		world.hp[id] = hp
		damage(world, id, dmg)
	}
}
