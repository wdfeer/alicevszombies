package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func newCombatText(world *World, position rl.Vector2, text string) Entity {
	id := world.newEntity()
	world.position[id] = position
	world.velocity[id] = rl.Vector2{X: 0, Y: -10}
	world.drag[id] = 0.5
	world.combatText[id] = text
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
	for id, str := range world.combatText {
		pos := world.position[id]

		color := rl.ColorAlpha(rl.White, min(rl.Vector2Length(world.velocity[id])/2-0.25, 1))
		rl.DrawTextEx(rl.GetFontDefault(), str, pos, 4, 1, color)
	}
}
