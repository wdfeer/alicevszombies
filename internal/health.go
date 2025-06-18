package internal

import (
	"fmt"
	"math/rand"

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
		rl.PlaySound(assets.sounds["player_hit"])
		world.combatText[ctextID] = CombatText{
			text: world.combatText[ctextID].text,
			hue:  rl.Red,
		}
	} else if _, ok := world.enemy[id]; ok {
		dist := rl.Vector2Distance(world.position[world.player], world.position[id])
		if dist < 200 {
			rl.SetSoundPitch(assets.sounds["enemy_hit"], 0.8+0.2*rand.Float32())
			rl.SetSoundVolume(assets.sounds["enemy_hit"], (1 - dist/200))
			rl.PlaySound(assets.sounds["enemy_hit"])
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
