package internal

type Upgrade = uint8

const (
	DOLL_DAMAGE Upgrade = iota
)

func incrementUpgrade(world *World, upgrade Upgrade) {
	lvl, exists := world.playerData.upgrades[upgrade]
	if exists {
		world.playerData.upgrades[upgrade] = lvl + 1
	} else {
		world.playerData.upgrades[upgrade] = 1
	}
}
