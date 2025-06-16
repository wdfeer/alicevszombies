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

		for doll := range world.dollTag {
			dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
			if rl.CheckCollisionRecs(dollRec, enemyRec) {
				// TODO: implement doll types

				baseDamage := 1
				if world.flipping[doll].baseTexture == "doll_lance" {
					baseDamage = 2
				}
				damageWithCooldown(world, enemy, float32(baseDamage)+(float32(world.playerData.upgrades[DOLL_DAMAGE])/4), doll)
				break
			}
		}
	}
}
