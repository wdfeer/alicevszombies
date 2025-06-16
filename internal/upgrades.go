package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE = "Doll Damage"
	DOLL_SPEED  = "Doll Speed"
	LANCE_DOLL  = "Lance Doll"
)

var allUpgrades = []Upgrade{DOLL_DAMAGE, DOLL_SPEED, LANCE_DOLL}

func getAvailableUpgrades(world *World) []Upgrade {
	newSlice := []Upgrade{}
	for _, up := range allUpgrades {
		if up == LANCE_DOLL {
			for _, typ := range world.doll {
				if typ == dollTypes.knifeDoll {
					newSlice = append(newSlice, up)
					break
				}
			}
		} else {
			newSlice = append(newSlice, up)
		}
	}
	return newSlice
}

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
	case LANCE_DOLL:
		sacrificed := false
		for id, typ := range world.doll {
			if typ == dollTypes.knifeDoll {
				sacrificed = true
				world.deleteEntity(id)
				break
			}
		}
		if sacrificed {
			id := newDoll(world, dollTypes.lanceDoll)
			world.position[id] = world.position[world.player]
		}
	}
}
