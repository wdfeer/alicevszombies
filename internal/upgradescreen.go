package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades [2]Upgrade
}

func newUpgradeScreen(world *World) {
	world.uistate.isUpgradeScreen = true
	world.uistate.upgradeScreen = UpgradeScreen{
		upgrades: randomUpgrades(world),
	}
}

func updateUpgradeScreen(world *World) {
	upgradeOne := rl.IsKeyPressed(rl.KeyOne)
	upgradeTwo := rl.IsKeyPressed(rl.KeyTwo)
	if upgradeOne || upgradeTwo {
		if upgradeOne {
			incrementUpgrade(world, world.uistate.upgradeScreen.upgrades[0])
		} else if upgradeTwo {
			incrementUpgrade(world, world.uistate.upgradeScreen.upgrades[1])
		}
		world.paused = false
		world.uistate.isUpgradeScreen = false
	}
}

func renderUpgradeScreen(world *World) {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))
	center := util.HalfScreenSize()
	util.DrawTextCenteredSpaced(world.uistate.upgradeScreen.upgrades[0], 40, rl.Vector2Add(center, rl.Vector2{X: -250, Y: -32}), 4)
	util.DrawTextCenteredSpaced("1", 64, rl.Vector2Add(center, rl.Vector2{X: -250, Y: 32}), 4)
	util.DrawTextCenteredSpaced(world.uistate.upgradeScreen.upgrades[1], 40, rl.Vector2Add(center, rl.Vector2{X: 250, Y: -32}), 4)
	util.DrawTextCenteredSpaced("2", 64, rl.Vector2Add(center, rl.Vector2{X: 250, Y: 32}), 4)
}
