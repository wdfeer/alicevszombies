package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIState struct {
	isMainMenu       bool
	isUpgradeScreen  bool
	isDeathScreen    bool
	previousMousePos rl.Vector2
	cursorHideTimer  float32
	upgradeScreen    UpgradeScreen
	mainMenu         MainMenu
}

func updateUI(world *World) {
	if world.uistate.isMainMenu {
	} else if world.uistate.isDeathScreen {
		updateDeathScreen(world)
	} else {
		if world.uistate.isUpgradeScreen {
			updateUpgradeScreen(world)
		} else if rl.IsKeyPressed(rl.KeyEscape) {
			world.paused = !world.paused
		}

		if rl.GetMousePosition() == world.uistate.previousMousePos {
			world.uistate.cursorHideTimer += dt
		} else {
			world.uistate.cursorHideTimer = 0
		}
		world.uistate.previousMousePos = rl.GetMousePosition()
	}
}

func renderHUD(world *World) {
	{ // Wave counter
		str := "Wave " + fmt.Sprint(world.enemySpawner.wave)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	{ // HP bar
		str := "HP: " + fmt.Sprint(world.hp[world.player].val)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) - 250}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	{ // MP bar
		str := "MP: " + fmt.Sprint(world.playerData.mana)
		center := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) - 200}
		pos := util.CenterText(str, 32, center)
		rl.DrawText(str, int32(pos.X), int32(pos.Y), 32, rl.White)
	}

	{ // Spell bar
		util.DrawTextureCenteredScaled(assets.textures["heal_icon"],
			rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 - 80},
			4)
		util.DrawTextCentered("H", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight())/2 - 80})
		util.DrawTextCentered("5 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 - 50})
		util.DrawTextureCenteredScaled(assets.textures["doll_icon"],
			rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight()) / 2},
			4)
		util.DrawTextCentered("J", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight()) / 2})
		util.DrawTextCentered("10 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 30})
		util.DrawTextureCenteredScaled(assets.textures["pitem_icon"],
			rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 80},
			4)
		util.DrawTextCentered("K", 40, rl.Vector2{X: 250, Y: float32(rl.GetScreenHeight())/2 + 80})
		util.DrawTextCentered("10 MP", 20, rl.Vector2{X: 200, Y: float32(rl.GetScreenHeight())/2 + 110})
	}
}

func renderUI(world *World) {
	if world.uistate.isMainMenu {
		renderMainMenu(world)
	} else if world.uistate.isDeathScreen {
		renderDeathScreen(world)
	} else {
		renderHUD(world)
		if world.uistate.isUpgradeScreen {
			renderUpgradeScreen(world)
		} else if world.paused {
			renderPauseMenu(world)
		}
	}

	if world.uistate.cursorHideTimer < 2.5 {
		rl.DrawTextureEx(assets.textures["cursor"+fmt.Sprint(world.uistate.cursorType)], rl.GetMousePosition(), 0, 4, rl.White)
	}
}
