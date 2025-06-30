package internal

import "math/rand"

type Upgrade struct {
	name         string
	dollType     *DollType
	cost         map[*DollType]uint8
	unique       bool
	incompatible []*Upgrade
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
	UpgradeSelection = Upgrade{
		name:   "Upgrade Selection",
		unique: true,
	}
	MovementSpeed = Upgrade{
		name:   "Move Speed",
		unique: true,
	}
	SprintUpgrade = Upgrade{
		name:   "Sprinting",
		unique: true,
	}
)

var uniqueUpgrades = []*Upgrade{&MovementSpeed, &UpgradeSelection}

func randomUpgradesFrom(world *World, available []*Upgrade) []*Upgrade {
	count := 2 + world.playerData.upgrades[&UpgradeSelection]
	upgrades := make([]*Upgrade, 0)
	for range count {
		if len(available) == 0 {
			break
		}
		index := rand.Int() % len(available)
		upgrade := available[index]
		unique := true
		for _, up := range upgrades {
			if up == upgrade {
				unique = false
			}
		}
		if unique {
			upgrades = append(upgrades, upgrade)
			available = append(available[:index], available[index+1:]...)
		}
	}
	return upgrades
}

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
				break
			}
		}

		if !failed && up.incompatible != nil {
			for _, x := range up.incompatible {
				if world.playerData.upgrades[x] > 0 {
					failed = true
					break
				}
			}
		}

		if !failed {
			newSlice = append(newSlice, up)
		}
	}
	return newSlice
}

func availableUniqueUpgrades(world *World) []*Upgrade {
	newSlice := make([]*Upgrade, 0)

	for _, up := range uniqueUpgrades {
		compatible := true
		if up.incompatible != nil {
			for _, x := range up.incompatible {
				if world.playerData.upgrades[x] > 0 {
					compatible = false
					break
				}
			}
		}
		if compatible && world.playerData.upgrades[up] == 0 {
			newSlice = append(newSlice, up)
		}
	}

	return newSlice
}

func randomUpgrades(world *World) []*Upgrade {
	return randomUpgradesFrom(world, availableUpgrades(world))
}

func randomUniqueUpgrades(world *World) []*Upgrade {
	return randomUpgradesFrom(world, availableUniqueUpgrades(world))
}

func incrementUpgrade(world *World, upgrade *Upgrade) {
	world.playerData.upgrades[upgrade]++

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
