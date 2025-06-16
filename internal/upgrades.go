package internal

import "math/rand"

type Upgrade = string

const (
	DOLL_DAMAGE = "Doll Damage"
	DOLL_SPEED  = "Doll Speed"
)

func randomUpgrade() Upgrade {
	upgrades := []Upgrade{DOLL_DAMAGE, DOLL_SPEED}
	return upgrades[rand.Int()%len(upgrades)]
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
