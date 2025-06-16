package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE = "Doll Damage"
	DOLL_SPEED  = "Doll Speed"
	FUSE_LANCE  = "Lance Doll"
)

var allUpgrades = []Upgrade{DOLL_DAMAGE, DOLL_SPEED, FUSE_LANCE}

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

	onUpgradeGet(world, upgrade)
}

func onUpgradeGet(world *World, upgrade Upgrade) {
	switch upgrade {
	case FUSE_LANCE:
		sacrificed := false
		for id, ok := range world.dollTag {
			if ok {
				sacrificed = true
				world.deleteEntity(id)
			}
		}
		if sacrificed {
			newLance(world)
		}
	}
}
