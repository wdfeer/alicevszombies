package internal

import (
	"alicevszombies/internal/util"
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades []*Upgrade
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
	oldFontSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 40)

	screen := world.uistate.upgradeScreen

	width := float32(440)

	center := util.HalfScreenSize()
	xPositions := util.SpaceCentered(width+120, len(screen.upgrades), center.X-width/2)
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	keys := map[int]int32{
		0: rl.KeyOne,
		1: rl.KeyTwo,
		2: rl.KeyThree,
	}
	upgrade := -1
	for i := range len(screen.upgrades) {
		rect := rl.NewRectangle(xPositions[i], center.Y-64, width, 128)
		raygui.Panel(rect, "")
		raygui.Label(rect, world.uistate.upgradeScreen.upgrades[i].name)
		rect.Y += 144
		rect.Height = 64
		if raygui.Button(rect, fmt.Sprint(i+1)) || rl.IsKeyPressed(keys[i]) {
			upgrade = i
		}
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontSize)

	if upgrade != -1 {
		incrementUpgrade(world, world.uistate.upgradeScreen.upgrades[upgrade])
		world.paused = false
		world.uistate.isUpgradeScreen = false
	}
}
