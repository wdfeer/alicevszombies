package internal

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIState struct {
	isMainMenu       bool
	isUpgradeScreen  bool
	isDeathScreen    bool
	previousMousePos rl.Vector2
	cursorHideTimer  float32
	pauseMenu        PauseMenu
	upgradeScreen    UpgradeScreen
	mainMenu         MainMenu
}

func updateUI(world *World) {
	if !rl.IsWindowFocused() {
		world.paused = true
	}

	if world.uistate.isMainMenu {
	} else if world.uistate.isDeathScreen {
		updateDeathScreen(world)
	} else {
		if rl.IsKeyPressed(rl.KeyEscape) && !world.uistate.isUpgradeScreen {
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
		rl.DrawTextureEx(assets.textures["cursor"+fmt.Sprint(options.CursorType)], rl.GetMousePosition(), 0, 4, rl.White)
	}
}
