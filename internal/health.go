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
		for id := range hp.attackerCooldown {
			hp.attackerCooldown[id] -= dt
		}

		world.hp[id] = HP{
			val:              hp.val,
			attackerCooldown: hp.attackerCooldown,
			immuneTime:       hp.immuneTime,
		}
	}
}

func heal(world *World, id Entity, amount float32) {
	hp := world.hp[id]
	hp.val += amount

	ctextID := newCombatText(world, world.position[id], fmt.Sprint(amount))
	world.combatText[ctextID] = CombatText{
		text: world.combatText[ctextID].text,
		hue:  rl.Green,
	}

	world.hp[id] = hp
}

func damage(world *World, id Entity, dmg float32) {
	hp, exists := world.hp[id]
	if !exists {
		println("WARNING: Tried damaging deleted enemy with id", id)
		return
	}

	hp.val -= dmg

	ctextID := newCombatText(world, world.position[id], fmt.Sprint(dmg))
	if dmg < 0 {
		world.combatText[ctextID] = CombatText{
			text: world.combatText[ctextID].text,
			hue:  rl.Green,
		}
	} else if id == world.player {
		world.combatText[ctextID] = CombatText{
			text: world.combatText[ctextID].text,
			hue:  rl.Red,
		}
	}

	if hp.val <= 0 {
		println("INFO: Entity with id", id, "killed!")
		world.deleteEntity(id)
	} else {
		world.hp[id] = hp
	}
}

func damageWithCooldown(world *World, id Entity, dmg float32, attacker Entity) {
	hp, exists := world.hp[id]
	if !exists {
		println("WARNING: Tried damaging deleted enemy with id", id)
		return
	}
	if cooldown, exists := hp.attackerCooldown[attacker]; !exists || cooldown <= 0 {
		hp.attackerCooldown[attacker] = hp.immuneTime
		world.hp[id] = hp
		damage(world, id, dmg)
	}
}

// Damages the enemy whilst applying upgrades
func damageEnemy(world *World, enemy Entity, baseDamage float32, source Entity) {
	dmg := baseDamage
	if _, ok := world.doll[source]; ok {
		dmg += float32(world.playerData.upgrades[DOLL_DAMAGE]) / 4
		damageWithCooldown(world, enemy, dmg, source)
	} else if proj, ok := world.projectile[source]; ok {
		dmg += float32(world.playerData.upgrades[DOLL_DAMAGE]) / 8
		if proj.typ.deleteOnHit {
			damage(world, enemy, dmg)
		} else {
			damageWithCooldown(world, enemy, dmg, source)
		}
	}
}
