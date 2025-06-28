package internal

func superUpgradesAvailable(world *World) bool {
	count := 0
	for _, v := range world.playerData.upgrades {
		count += int(v)
	}
	return count > 3
}

func newSuperUpgradeScreen(world *World) {
	// TODO: implement
}
