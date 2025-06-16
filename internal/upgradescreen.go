package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades [2]Upgrade
}

func newUpgradeScreen(world *World) {
	uistate.upgradeScreenShown = true
	uistate.upgradeScreen = UpgradeScreen{
		upgrades: [2]Upgrade{DOLL_DAMAGE, DOLL_SPEED},
	}
}

func updateUpgradeScreen(world *World) {
	upgradeOne := rl.IsKeyPressed(rl.KeyOne)
	upgradeTwo := rl.IsKeyPressed(rl.KeyTwo)
	if upgradeOne || upgradeTwo {
		if upgradeOne {
			incrementUpgrade(world, DOLL_DAMAGE)
		} else if upgradeTwo {
			incrementUpgrade(world, DOLL_SPEED)
		}
		world.paused = false
		uistate.upgradeScreenShown = false
	}
}

func renderUpgradeScreen(world *World) {
	center := util.GetHalfScreen()
	util.DrawTextCenteredSpaced("Increase Doll Damage", 40, rl.Vector2Add(center, rl.Vector2{X: -250, Y: -32}), 4)
	util.DrawTextCenteredSpaced("1", 64, rl.Vector2Add(center, rl.Vector2{X: -250, Y: 32}), 4)
	util.DrawTextCenteredSpaced("Increase Doll Speed", 40, rl.Vector2Add(center, rl.Vector2{X: 250, Y: -32}), 4)
	util.DrawTextCenteredSpaced("2", 64, rl.Vector2Add(center, rl.Vector2{X: 250, Y: 32}), 4)
}
