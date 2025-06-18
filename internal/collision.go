package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateCollisions(world *World) {
	playerRec := util.CenterRectangle(world.position[world.player], world.size[world.player])

	for enemy := range world.enemy {
		enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])

		// Enemy -> Player
		if rl.CheckCollisionRecs(playerRec, enemyRec) {
			damageWithCooldown(world, world.player, 1, enemy)
		}

		// Doll -> Enemy
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

		// Projectile -> Enemy
		for id, proj := range world.projectile {
			if proj.typ.hostile {
				continue
			}

			projRec := util.CenterRectangle(world.position[id], proj.typ.size)
			if rl.CheckCollisionRecs(enemyRec, projRec) {
				if proj.typ.deleteOnHit {
					damage(world, enemy, proj.typ.damage+(float32(world.playerData.upgrades[DOLL_DAMAGE])/8))
					world.deleteEntity(id)
				} else {
					damageWithCooldown(world, enemy, proj.typ.damage+(float32(world.playerData.upgrades[DOLL_DAMAGE])/8), id)
				}
				break
			}
		}
	}

	// Projectile -> Player
	for id, proj := range world.projectile {
		if !proj.typ.hostile {
			continue
		}

		projRec := util.CenterRectangle(world.position[id], proj.typ.size)
		if rl.CheckCollisionRecs(playerRec, projRec) {

			damageWithCooldown(world, world.player, proj.typ.damage, id)
			break
		}
	}
}
