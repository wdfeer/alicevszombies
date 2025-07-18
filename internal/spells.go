package internal

import (
	"alicevszombies/internal/util"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderSpells(world *World) {
	halfHeight := float32(rl.GetScreenHeight() / 2)
	size := rl.Vector2{X: 200, Y: 80}
	pos := rl.Vector2{X: 300, Y: halfHeight}

	spellCount := 3
	if world.playerData.upgradeCount() > 10 {
		spellCount = 4
	}
	yPositions := util.SpaceCentered(size.Y*1.2, spellCount, pos.Y)

	// Heal Spell
	pos.Y = yPositions[0]
	canHeal := world.playerData.mana >= 5 && !(world.difficulty == LUNATIC && world.status[world.player][Poison] > 0)
	if !canHeal {
		raygui.Disable()
	}
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyOne)) && canHeal && !world.paused {
		heal(world, world.player, 5)
		world.playerData.mana -= 5
	}
	raygui.Enable()
	util.DrawTextureCenteredScaled(assets.textures["heal_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("1", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	// Doll Summon
	if world.playerData.mana < 10 || world.playerData.dollToSpawn != nil {
		raygui.Disable()
	}
	pos.Y = yPositions[1]
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyTwo)) && (world.playerData.mana >= 10 && world.playerData.dollToSpawn == nil) && !world.paused {
		spawnDollWithAnimation(world, &dollTypes.basicDoll)
		world.playerData.mana -= 10
	}
	util.DrawTextureCenteredScaled(assets.textures["doll_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("2", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	// Upgrade
	if world.playerData.mana >= 10 {
		raygui.Enable()
	}
	pos.Y = yPositions[2]
	if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyThree)) && world.playerData.mana >= 10 && !world.paused {
		world.paused = true
		world.playerData.mana -= 10
		newUpgradeScreen(world)
	}
	util.DrawTextureCenteredScaled(assets.textures["pitem_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
	util.DrawTextCentered("3", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})

	// Unique Upgrade
	uniqueUpgradeCost := float32(50)
	for _, up := range uniqueUpgrades {
		uniqueUpgradeCost += float32(world.playerData.upgrades[up]) * 25
	}
	if spellCount == 4 {
		disable := world.playerData.mana < uniqueUpgradeCost || len(availableUniqueUpgrades(world)) == 0
		if disable {
			raygui.Disable()
		}
		pos.Y = yPositions[3]
		if (raygui.Button(util.CenterRectangle(pos, size), "") || rl.IsKeyPressed(rl.KeyFour)) && !disable && !world.paused {
			world.paused = true
			world.playerData.mana -= uniqueUpgradeCost
			newUniqueUpgradeScreen(world)
		}
		util.DrawTextureCenteredScaled(assets.textures["unique_upgrade_icon"], rl.Vector2{X: pos.X - size.X/5, Y: pos.Y}, 4)
		util.DrawTextCentered("4", 40, rl.Vector2{X: pos.X + size.X/5, Y: pos.Y})
	}

	raygui.Enable()
}
