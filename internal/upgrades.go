package internal

import "math/rand"

type Upgrade struct {
	name     string
	dollType *DollType
	cost     map[*DollType]uint8
}

var (
	DollDamage = Upgrade{
		name: "Doll Damage",
	}
	DollSpeed = Upgrade{
		name: "Doll Speed",
	}
	LanceDoll = Upgrade{
		name:     "Lance Doll",
		dollType: &dollTypes.lanceDoll,
		cost:     map[*DollType]uint8{&dollTypes.basicDoll: 1},
	}
	ScytheDoll = Upgrade{
		name:     "Scythe Doll",
		dollType: &dollTypes.scytheDoll,
		cost:     map[*DollType]uint8{&dollTypes.lanceDoll: 2},
	}
	KnifeDoll = Upgrade{
		name:     "Knife Doll",
		dollType: &dollTypes.knifeDoll,
		cost:     map[*DollType]uint8{&dollTypes.basicDoll: 1},
	}
	MagicianDoll = Upgrade{
		name:     "Magician Doll",
		dollType: &dollTypes.magicianDoll,
		cost:     map[*DollType]uint8{&dollTypes.knifeDoll: 2},
	}
)

var allUpgrades = []Upgrade{DollDamage, DollSpeed, LanceDoll, ScytheDoll, KnifeDoll, MagicianDoll}

func getAvailableUpgrades(world *World) []Upgrade {
	newSlice := []Upgrade{}

	dollCounts := make(map[*DollType]uint8, 0)

	for _, typ := range world.doll {
		if _, ok := dollCounts[typ]; ok {
			dollCounts[typ]++
		} else {
			dollCounts[typ] = 1
		}
	}

	for _, up := range allUpgrades {
		if up.cost == nil {
			newSlice = append(newSlice, up)
			continue
		}

		failed := false
		for doll, required := range up.cost {
			if count, ok := dollCounts[doll]; required > 0 && (!ok || count < required) {
				failed = true
			}
		}

		if !failed {
			newSlice = append(newSlice, up)
		}
	}
	return newSlice
}

func randomUpgrades(world *World) [2]*Upgrade {
	available := getAvailableUpgrades(world)
	upgrade1 := available[rand.Int()%len(available)]
	upgrade2 := upgrade1
	for upgrade2.name == upgrade1.name {
		upgrade2 = available[rand.Int()%len(available)]
	}
	return [2]*Upgrade{&upgrade1, &upgrade2}
}

func incrementUpgrade(world *World, upgrade *Upgrade) {
	lvl, exists := world.playerData.upgrades[upgrade]
	if exists {
		world.playerData.upgrades[upgrade] = lvl + 1
	} else {
		world.playerData.upgrades[upgrade] = 1
	}

	pos := world.position[world.player]
	pos.Y -= 5
	newCombatText(world, pos, upgrade.name+" +")

	onUpgradeGet(world, upgrade)
}

func onUpgradeGet(world *World, upgrade *Upgrade) {
	if upgrade.dollType == nil {
		return
	}

	for typ, count := range upgrade.cost {
		for _, d := range world.doll {
			if d == typ {
				count--
			}
			if count <= 0 {
				break
			}
		}
	}

	id := newDoll(world, upgrade.dollType)
	world.position[id] = world.position[world.player]
}
