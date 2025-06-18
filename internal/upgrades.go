package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE   = "Doll Damage"
	DOLL_SPEED    = "Doll Speed"
	LANCE_DOLL    = "Lance Doll"
	SCYTHE_DOLL   = "Scythe Doll"
	KNIFE_DOLL    = "Knife Doll"
	MAGICIAN_DOLL = "Magician Doll"
)

var allUpgrades = []Upgrade{DOLL_DAMAGE, DOLL_SPEED, LANCE_DOLL, SCYTHE_DOLL, KNIFE_DOLL, MAGICIAN_DOLL}

func getAvailableUpgrades(world *World) []Upgrade {
	newSlice := []Upgrade{}

	basicDollCount := 0
	lanceDollCount := 0
	knifeDollCount := 0
	for _, typ := range world.doll {
		switch typ {
		case &dollTypes.basicDoll:
			basicDollCount++
		case &dollTypes.lanceDoll:
			lanceDollCount++
		case &dollTypes.knifeDoll:
			knifeDollCount++
		}
	}

	for _, up := range allUpgrades {
		switch up {
		default:
			newSlice = append(newSlice, up)
		case DOLL_SPEED:
			if basicDollCount == 0 {

				newSlice = append(newSlice, up)
			} else {
				meleeCount := 0
				for _, typ := range world.doll {
					if typ.projectileType == nil {
						meleeCount++
						if meleeCount > 1 {
							newSlice = append(newSlice, up)
							break
						}
					}
				}
			}
		case LANCE_DOLL:
			fallthrough
		case KNIFE_DOLL:
			if basicDollCount > 0 {
				newSlice = append(newSlice, up)
			}
		case MAGICIAN_DOLL:
			if knifeDollCount > 1 {
				newSlice = append(newSlice, up)
			}
		case SCYTHE_DOLL:
			if lanceDollCount > 1 {
				newSlice = append(newSlice, up)
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
	// TODO: make Upgrade a struct to refactor this somehow
	dollUpgrades := map[Upgrade]*DollType{
		LANCE_DOLL:    &dollTypes.lanceDoll,
		SCYTHE_DOLL:   &dollTypes.scytheDoll,
		KNIFE_DOLL:    &dollTypes.knifeDoll,
		MAGICIAN_DOLL: &dollTypes.magicianDoll,
	}

	for up, dollType := range dollUpgrades {
		if up != upgrade {
			continue
		}

		for id, typ := range world.doll {
			if up != MAGICIAN_DOLL && up != SCYTHE_DOLL {
				if typ == &dollTypes.basicDoll {
					world.deleteEntity(id)
					break
				}
			} else {
				var desiredType *DollType
				if up == MAGICIAN_DOLL {
					desiredType = &dollTypes.knifeDoll
				} else {
					desiredType = &dollTypes.lanceDoll
				}
				count := 0
				if typ == desiredType {
					world.deleteEntity(id)
					count++
					if count >= 2 {
						break
					}
				}
			}
		}

		id := newDoll(world, dollType)
		world.position[id] = world.position[world.player]

		break
	}
}
