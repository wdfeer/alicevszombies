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
				damageEnemy(world, enemy, typ.contactDamage, doll)
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
					damageEnemy(world, enemy, proj.typ.damage, id)
					world.deleteEntity(id)
				} else {
					damageWithCooldown(world, enemy, proj.typ.damage, id)
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
