package internal

import "math/rand"

type Upgrade struct {
	name     string
	dollType *DollType
	cost     map[*DollType]uint8
	super    bool
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
	DestructionDoll = Upgrade{
		name:     "Destruction Doll",
		dollType: &dollTypes.destructionDoll,
		cost:     map[*DollType]uint8{&dollTypes.magicianDoll: 2, &dollTypes.scytheDoll: 2},
	}
)

var upgrades = []*Upgrade{&DollDamage, &DollSpeed, &LanceDoll, &ScytheDoll, &KnifeDoll, &MagicianDoll, &DestructionDoll}

var (
	MovementSpeed = Upgrade{
		name:  "Move Speed",
		super: true,
	}
)

var superUpgrades = []*Upgrade{&MovementSpeed}

func availableUpgrades(world *World) []*Upgrade {
	newSlice := []*Upgrade{}

	dollCounts := make(map[*DollType]uint8, 0)

	for _, typ := range world.doll {
		dollCounts[typ]++
	}

	for _, up := range upgrades {
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

func randomUpgrades(world *World) []*Upgrade {
	available := availableUpgrades(world)
	upgrade1 := available[rand.Int()%len(available)]
	upgrade2 := upgrade1
	for upgrade2.name == upgrade1.name {
		upgrade2 = available[rand.Int()%len(available)]
	}
	return []*Upgrade{upgrade1, upgrade2}
}

func availableSuperUpgrades(world *World) []*Upgrade {
	newSlice := make([]*Upgrade, 0)

	for _, up := range superUpgrades {
		if world.playerData.upgrades[up] == 0 {
			newSlice = append(newSlice, up)
		}
	}

	return newSlice
}

func randomSuperUpgrades(world *World) []*Upgrade {
	available := availableSuperUpgrades(world)
	upgrade1 := available[rand.Int()%len(available)]
	upgrade2 := available[rand.Int()%len(available)]
	if upgrade1 != upgrade2 {
		return []*Upgrade{upgrade1, upgrade2}
	} else {
		return []*Upgrade{upgrade1}
	}
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
		for id, d := range world.doll {
			if d == typ {
				count--
				world.deleteEntity(id)
			}
			if count <= 0 {
				break
			}
		}
	}

	id := newDoll(world, upgrade.dollType)
	world.position[id] = world.position[world.player]
}
