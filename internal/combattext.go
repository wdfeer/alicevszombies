package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type CombatText struct {
	text string
	hue  rl.Color
}

func newCombatText(world *World, position rl.Vector2, text string) Entity {
	id := world.newEntity()
	world.position[id] = position
	world.velocity[id] = rl.Vector2{X: 0, Y: -10}
	world.drag[id] = 0.5
	world.combatText[id] = CombatText{
		text: text,
		hue:  rl.White,
	}
	return id
}

func updateCombatText(world *World) {
	for id := range world.combatText {
		if rl.Vector2Length(world.velocity[id]) < 0.5 {
			world.deleteEntity(id)
		}
	}
}

func renderCombatText(world *World) {
	for id, ctext := range world.combatText {
		pos := world.position[id]

		color := rl.ColorAlpha(ctext.hue, min(rl.Vector2Length(world.velocity[id])/2-0.25, 1))
		rl.DrawTextEx(rl.GetFontDefault(), ctext.text, pos, 4, 1, color)
	}
}
