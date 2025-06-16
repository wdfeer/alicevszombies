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

		for doll, typ := range world.doll {
			if typ.contactDamage <= 0 {
				continue
			}

			dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
			if rl.CheckCollisionRecs(dollRec, enemyRec) {
				damageWithCooldown(world, enemy, typ.contactDamage+(float32(world.playerData.upgrades[DOLL_DAMAGE])/4), doll)
				break
			}
		}

		for id, proj := range world.projectile {
			projRec := util.CenterRectangle(world.position[id], proj.typ.size)
			if rl.CheckCollisionRecs(enemyRec, projRec) {
				if proj.typ.deleteOnHit {
					damage(world, enemy, proj.typ.damage)
					world.deleteEntity(id)
				} else {
					damageWithCooldown(world, enemy, proj.typ.damage, id)
				}
				break
			}
		}
	}
}
