package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE   = "Doll Damage"
	DOLL_SPEED    = "Doll Speed"
	LANCE_DOLL    = "Lance Doll"
	KNIFE_DOLL    = "Knife Doll"
	MAGICIAN_DOLL = "Magician Doll"
)

var allUpgrades = []Upgrade{DOLL_DAMAGE, DOLL_SPEED, LANCE_DOLL, KNIFE_DOLL, MAGICIAN_DOLL}

func getAvailableUpgrades(world *World) []Upgrade {
	newSlice := []Upgrade{}
	for _, up := range allUpgrades {
		switch up {
		default:
			newSlice = append(newSlice, up)
		case LANCE_DOLL:
			fallthrough
		case KNIFE_DOLL:
			fallthrough
		case MAGICIAN_DOLL:
			for _, typ := range world.doll {
				if typ == &dollTypes.swordDoll {
					newSlice = append(newSlice, up)
					break
				}
			}
		}
	}
	return newSlice
}

func randomUpgrades(world *World) [2]Upgrade {
	available := getAvailableUpgrades(world)
	upgrade1 := available[rand.Int()%len(available)]
	upgrade2 := upgrade1
	for upgrade2 == upgrade1 {
		upgrade2 = available[rand.Int()%len(available)]
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
	dollUpgrades := map[Upgrade]*DollType{
		LANCE_DOLL:    &dollTypes.lanceDoll,
		KNIFE_DOLL:    &dollTypes.knifeDoll,
		MAGICIAN_DOLL: &dollTypes.magicianDoll,
	}

	for up, dollType := range dollUpgrades {
		if up == upgrade {
			sacrificed := false
			for id, typ := range world.doll {
				if typ == &dollTypes.swordDoll {
					sacrificed = true
					world.deleteEntity(id)
					break
				}
			}
			if sacrificed {
				id := newDoll(world, dollType)
				world.position[id] = world.position[world.player]
			}

			break
		}
	}
}
