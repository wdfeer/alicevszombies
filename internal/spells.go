package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateSpells(world *World) {
	if world.playerData.mana >= 5 && (rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyH)) {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyJ)) {
		id := newDoll(world, &dollTypes.basicDoll)
		world.position[id] = world.position[world.player]
		world.playerData.mana -= 10
	}
	if world.playerData.mana >= 10 && (rl.IsKeyPressed(rl.KeyThree) || rl.IsKeyPressed(rl.KeyK)) {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
}

func renderSpellsBar() {
	screenHeight := float32(rl.GetScreenHeight())

	drawCenteredImage := func(tex rl.Texture2D, center rl.Vector2, scale float32) {
		width := float32(tex.Width) * scale
		height := float32(tex.Height) * scale
		pos := rl.NewRectangle(center.X-width/2, center.Y-height/2, width, height)
		rl.DrawTexturePro(tex, rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)), pos, rl.Vector2{}, 0, rl.White)
	}

	drawCenteredText := func(text string, fontSize int, center rl.Vector2) {
		textWidth := float32(rl.MeasureText(text, int32(fontSize)))
		rl.DrawText(text, int32(center.X-textWidth/2), int32(center.Y-float32(fontSize)/2), int32(fontSize), rl.White)
	}

	drawCenteredImage(assets.textures["heal_icon"], rl.NewVector2(200, screenHeight/2-80), 4)
	drawCenteredText("H", 40, rl.NewVector2(250, screenHeight/2-80))
	drawCenteredText("5 MP", 20, rl.NewVector2(200, screenHeight/2-50))

	drawCenteredImage(assets.textures["doll_icon"], rl.NewVector2(200, screenHeight/2), 4)
	drawCenteredText("J", 40, rl.NewVector2(250, screenHeight/2))
	drawCenteredText("10 MP", 20, rl.NewVector2(200, screenHeight/2+30))

	drawCenteredImage(assets.textures["pitem_icon"], rl.NewVector2(200, screenHeight/2+80), 4)
	drawCenteredText("K", 40, rl.NewVector2(250, screenHeight/2+80))
	drawCenteredText("10 MP", 20, rl.NewVector2(200, screenHeight/2+110))
}
