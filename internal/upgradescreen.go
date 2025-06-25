package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades [2]*Upgrade
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
	oldFontSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 40)

	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))
	center := util.HalfScreenSize()
	rect := rl.NewRectangle(center.X-320-60, center.Y-64, 320, 128)
	raygui.Panel(rect, "")
	raygui.Label(rect, world.uistate.upgradeScreen.upgrades[0].name)
	rect.Y += 144
	rect.Height = 64
	raygui.Button(rect, "1")
	rect.X += 320 + 60*2
	raygui.Button(rect, "2")
	rect.Y -= 144
	rect.Height = 128
	raygui.Panel(rect, "")
	raygui.Label(rect, world.uistate.upgradeScreen.upgrades[1].name)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontSize)
}
