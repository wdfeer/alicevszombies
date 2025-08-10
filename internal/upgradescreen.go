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

	chosenUpgrade          *Upgrade
	chosenUpgradeIndex     int
	chosenUpgradeAnimTimer float32
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

	screen := &world.uistate.upgradeScreen

	width := float32(480) * uiScale
	height := float32(240) * uiScale

	center := util.HalfScreenSize()
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.4))

	{
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize64)
		raygui.Label(rl.Rectangle{X: center.X - 320*uiScale, Y: center.Y - height*1.6, Width: 640 * uiScale, Height: 120 * uiScale}, "Select Upgrade")
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)
	}

	keys := map[int]int32{
		0: rl.KeyOne,
		1: rl.KeyTwo,
		2: rl.KeyThree,
	}
	upgrade := -1
	xPositions := util.SpaceCentered(width+120*uiScale, len(screen.upgrades), center.X-width/2)
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
			dollCostRect.X += 32 * uiScale
			dollCostRect.Width -= 64 * uiScale
			dollCostRect.Y += height / 8
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
			raygui.Label(dollCostRect, "Cost:")
			raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_CENTER))

			// Calculate dimensions of the doll cost
			dollScale := 4 * uiScale
			dollSpacing := 4 * uiScale
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
					x -= float32(assets.textures[data.typ.texture].Width*int32(dollScale) + int32(dollSpacing))
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
		upgradeType := world.uistate.upgradeScreen.upgrades[upgrade]

		incrementUpgrade(world, upgradeType)
		world.paused = false
		world.uistate.isUpgradeScreen = false

		screen.chosenUpgrade = upgradeType
		screen.chosenUpgradeIndex = upgrade
		screen.chosenUpgradeAnimTimer = 1
	}

	world.uistate.upgradeScreen.active = true
}

func renderChosenUpgradeAnimation(world *World) { // TODO: refactor this shit
	screen := &world.uistate.upgradeScreen
	if screen.chosenUpgradeAnimTimer <= 0 {
		return
	}

	raygui.SetState(raygui.STATE_PRESSED)

	oldColor1 := raygui.GetStyle(raygui.DEFAULT, raygui.BACKGROUND_COLOR)
	newColor1 := rl.ColorAlpha(rl.GetColor(uint(oldColor1)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.BACKGROUND_COLOR, int64(rl.ColorToInt(newColor1)))
	oldColor2 := raygui.GetStyle(raygui.DEFAULT, raygui.BASE_COLOR_PRESSED)
	newColor2 := rl.ColorAlpha(rl.GetColor(uint(oldColor2)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.BASE_COLOR_PRESSED, int64(rl.ColorToInt(newColor2)))
	oldColor3 := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_PRESSED)
	newColor3 := rl.ColorAlpha(rl.GetColor(uint(oldColor3)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_PRESSED, int64(rl.ColorToInt(newColor3)))
	oldColor4 := raygui.GetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_PRESSED)
	newColor4 := rl.ColorAlpha(rl.GetColor(uint(oldColor4)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_PRESSED, int64(rl.ColorToInt(newColor4)))
	oldColor5 := raygui.GetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_NORMAL)
	newColor5 := rl.ColorAlpha(rl.GetColor(uint(oldColor5)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_NORMAL, int64(rl.ColorToInt(newColor5)))
	oldColor6 := raygui.GetStyle(raygui.DEFAULT, raygui.LINE_COLOR)
	newColor6 := rl.ColorAlpha(rl.GetColor(uint(oldColor6)), screen.chosenUpgradeAnimTimer)
	raygui.SetStyle(raygui.DEFAULT, raygui.LINE_COLOR, int64(rl.ColorToInt(newColor6)))

	oldFontSize := raygui.GetStyle(raygui.DEFAULT, raygui.TEXT_SIZE)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize40)

	width := float32(480) * uiScale
	height := float32(240) * uiScale

	center := util.HalfScreenSize()

	i := screen.chosenUpgradeIndex
	up := screen.chosenUpgrade
	xPositions := util.SpaceCentered(width+120*uiScale, len(screen.upgrades), center.X-width/2)

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
		dollCostRect.X += 32 * uiScale
		dollCostRect.Width -= 64 * uiScale
		dollCostRect.Y += height / 8
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_LEFT))
		raygui.Label(dollCostRect, "Cost:")
		raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, int64(raygui.TEXT_ALIGN_CENTER))

		// Calculate dimensions of the doll cost
		dollScale := 4 * uiScale
		dollSpacing := 4 * uiScale
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
				rl.DrawTextureEx(assets.textures[data.typ.texture], pos, 0, float32(dollScale), rl.ColorAlpha(rl.White, screen.chosenUpgradeAnimTimer))
				x -= float32(assets.textures[data.typ.texture].Width*int32(dollScale) + int32(dollSpacing))
			}
		}
	}

	rect.Y += height + 48
	rect.Height = 64
	raygui.Button(rect, fmt.Sprint(i+1))

	screen.chosenUpgradeAnimTimer -= dt * 2

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, oldFontSize)
	raygui.SetStyle(raygui.DEFAULT, raygui.BACKGROUND_COLOR, oldColor1)
	raygui.SetStyle(raygui.DEFAULT, raygui.BASE_COLOR_PRESSED, oldColor2)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_PRESSED, oldColor3)
	raygui.SetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_PRESSED, oldColor4)
	raygui.SetStyle(raygui.DEFAULT, raygui.BORDER_COLOR_NORMAL, oldColor5)
	raygui.SetStyle(raygui.DEFAULT, raygui.LINE_COLOR, oldColor6)
	raygui.SetState(raygui.STATE_NORMAL)
}
