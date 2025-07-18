package internal

import (
	"alicevszombies/internal/util"
	"fmt"
	"sort"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UpgradeScreen struct {
	upgrades []*Upgrade
	// To prevent getting upgrade 3 right after pressing 3 to open the upgrade screen
	// set to false on first frame
	active bool
}

func newUpgradeScreen(world *World) {
	world.uistate.isUpgradeScreen = true
	world.uistate.upgradeScreen = UpgradeScreen{
		upgrades: randomUpgrades(world),
	}
}

func newUniqueUpgradeScreen(world *World) {
	world.uistate.isUpgradeScreen = true
	world.uistate.upgradeScreen = UpgradeScreen{
		upgrades: randomUniqueUpgrades(world),
	}
}

func renderUpgradeScreen(world *World) {
	oldFontSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 40)

	screen := world.uistate.upgradeScreen

	const width = float32(480)
	const height = float32(240)

	center := util.HalfScreenSize()
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	oldTextSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 64)
	raygui.Label(rl.Rectangle{X: center.X - 320, Y: center.Y - height*1.6, Width: 640, Height: 120}, "Select Upgrade")
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldTextSize)

	keys := map[int]int32{
		0: rl.KeyOne,
		1: rl.KeyTwo,
		2: rl.KeyThree,
	}
	upgrade := -1
	xPositions := util.SpaceCentered(width+120, len(screen.upgrades), center.X-width/2)
	for i, up := range screen.upgrades {
		rect := rl.Rectangle{X: xPositions[i], Y: center.Y - height/2, Width: width, Height: height}
		raygui.Panel(rect, "")

		titleRect := rect
		if up.cost == nil {
			titleRect.Y -= height / 8
		} else {
			titleRect.Y -= height / 3
		}
		raygui.Label(titleRect, world.uistate.upgradeScreen.upgrades[i].name)

		// Doll Cost
		if up.cost != nil {
			dollCostRect := rect
			dollCostRect.X += 32
			dollCostRect.Width -= 64
			dollCostRect.Y += height / 8
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
			raygui.Label(dollCostRect, "Cost:")
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_CENTER))

			// Calculate dimensions of the doll cost
			const dollScale = 4
			const dollSpacing = 4
			maxHeight := int32(0)
			for dollT := range up.cost {
				if assets.textures[dollT.texture].Height > maxHeight {
					maxHeight = assets.textures[dollT.texture].Height
				}
			}

			// Sort the doll costs by type
			type dollCost struct {
				typ   *DollType
				count uint8
			}
			sortedCosts := make([]dollCost, 0, len(up.cost))
			for typ, count := range up.cost {
				sortedCosts = append(sortedCosts, dollCost{typ, count})
			}
			sort.Slice(sortedCosts, func(i, j int) bool {
				return sortedCosts[i].typ.texture < sortedCosts[j].typ.texture
			})

			// Rendering of Dolls
			x := xPositions[i] + dollCostRect.Width
			for _, data := range sortedCosts {
				for range data.count {
					pos := rl.Vector2{X: x, Y: dollCostRect.Y + dollCostRect.Height/2 - float32(assets.textures[data.typ.texture].Height)*dollScale/2}
					rl.DrawTextureEx(assets.textures[data.typ.texture], pos, 0, float32(dollScale), rl.White)
					x -= float32(assets.textures[data.typ.texture].Width*dollScale + dollSpacing)
				}
			}
		}

		rect.Y += height + 48
		rect.Height = 64
		if raygui.Button(rect, fmt.Sprint(i+1)) || rl.IsKeyPressed(keys[i]) && world.uistate.upgradeScreen.active {
			upgrade = i
		}
	}

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontSize)

	if upgrade != -1 {
		incrementUpgrade(world, world.uistate.upgradeScreen.upgrades[upgrade])
		world.paused = false
		world.uistate.isUpgradeScreen = false
	}

	world.uistate.upgradeScreen.active = true
}
