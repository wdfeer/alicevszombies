package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateCollisions(world *World) {
	playerRec := util.CenterRectangle(world.position[world.player], world.size[world.player])

	for enemy := range world.enemyTag {
		enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])

		if rl.CheckCollisionRecs(playerRec, enemyRec) {
			damageWithCooldown(world, world.player, 1, enemy)
		}

		for dollID, dollTyp := range world.doll {
			dollRec := util.CenterRectangle(world.position[dollID], world.size[dollID])
			if rl.CheckCollisionRecs(dollRec, enemyRec) {
				damageWithCooldown(world, enemy, dollTyp.baseDamage+(float32(world.playerData.upgrades[DOLL_DAMAGE])/4), dollID)
				break
			}
		}
	}
}
