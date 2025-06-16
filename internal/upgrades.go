package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE = "Doll Damage"
	DOLL_SPEED  = "Doll Speed"
)

var allUpgrades = []Upgrade{DOLL_DAMAGE, DOLL_SPEED}

func randomUpgrades() [2]Upgrade {
	upgrade1 := allUpgrades[rand.Int()%len(allUpgrades)]
	upgrade2 := upgrade1
	for upgrade2 == upgrade1 {
		upgrade2 = allUpgrades[rand.Int()%len(allUpgrades)]
	}
	return [2]Upgrade{upgrade1, upgrade2}
}

func incrementUpgrade(world *World, upgrade Upgrade) {
	lvl, exists := world.playerData.upgrades[upgrade]
	if exists {
		world.playerData.upgrades[upgrade] = lvl + 1
	} else {
		world.playerData.upgrades[upgrade] = 1
	}

	pos := world.position[world.player]
	pos.Y -= 5
	newCombatText(world, pos, upgrade+" +")
}
