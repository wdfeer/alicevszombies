package internal

type Difficulty = uint8

const (
	EASY Difficulty = iota
	NORMAL
	HARD
	LUNATIC
)

func selectDifficulty(world *World, difficulty Difficulty) {
	world.difficulty = difficulty

	var iTime float32
	switch world.difficulty {
	case EASY:
		iTime = 1.8
	case NORMAL:
		iTime = 1.2
	case HARD:
		iTime = 1
	case LUNATIC:
		iTime = 0.8
	}
	world.hp[world.player] = HP{
		val:              10,
		immuneTime:       iTime,
		attackerCooldown: make(map[Entity]float32),
	}

	mana := float32(0)
	if world.difficulty == EASY {
		mana += 10
	}
	world.playerData = PlayerData{
		mana:     mana,
		upgrades: make(map[Upgrade]uint32),
	}

	stats.RunCount[difficulty]++
}
