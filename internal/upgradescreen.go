package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades [2]*Upgrade // TODO: unspecify size
}

func newUpgradeScreen(world *World) {
	world.uistate.isUpgradeScreen = true
	world.uistate.upgradeScreen = UpgradeScreen{
		upgrades: randomUpgrades(world),
	}
}

func newSuperUpgradeScreen(world *World) {
	world.uistate.isUpgradeScreen = true
	world.uistate.upgradeScreen = UpgradeScreen{
		upgrades: randomSuperUpgrades(world),
	}
}

func renderUpgradeScreen(world *World) {
	upgrade := -1

	oldFontSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 40)

	width := float32(440)

	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))
	center := util.HalfScreenSize()
	rect := rl.NewRectangle(center.X-width-60, center.Y-64, width, 128)
	raygui.Panel(rect, "")
	raygui.Label(rect, world.uistate.upgradeScreen.upgrades[0].name)
	rect.Y += 144
	rect.Height = 64
	if raygui.Button(rect, "1") || rl.IsKeyPressed(rl.KeyOne) {
		upgrade = 0
	}

	rect.X += width + 60*2
	if raygui.Button(rect, "2") || rl.IsKeyPressed(rl.KeyTwo) {
		upgrade = 1
	}
	rect.Y -= 144
	rect.Height = 128
	raygui.Panel(rect, "")
	raygui.Label(rect, world.uistate.upgradeScreen.upgrades[1].name)

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontSize)

	if upgrade != -1 {
		incrementUpgrade(world, world.uistate.upgradeScreen.upgrades[upgrade])
		world.paused = false
		world.uistate.isUpgradeScreen = false
	}
}
