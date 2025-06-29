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

			dir := util.Vector2Direction(world.position[world.player], world.position[enemy])
			world.velocity[enemy] = rl.Vector2Add(world.velocity[enemy], rl.Vector2Scale(dir, 800*dt))
		}

		// Doll -> Enemy
		for doll, typ := range world.doll {
			if typ.contactDamage <= 0 || typ.size.X <= 0 {
				continue
			}

			dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
			if rl.CheckCollisionRecs(dollRec, enemyRec) {
				if typ == &dollTypes.scytheDoll {
					damageWithCooldown(world, enemy, typ.contactDamage+(float32(world.playerData.upgrades[&DollDamage])/2), doll)
				} else {
					damageWithCooldown(world, enemy, typ.contactDamage+(float32(world.playerData.upgrades[&DollDamage])/4), doll)
				}
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
					damage(world, enemy, proj.typ.damage+(float32(world.playerData.upgrades[&DollDamage])/8))
					world.deleteEntity(id)
				} else {
					damageWithCooldown(world, enemy, proj.typ.damage+(float32(world.playerData.upgrades[&DollDamage])/8), id)
				}
				break
			}
		}

		// Enemy <-> Enemy, knockback only
		enemyRec.Width /= 2
		enemyRec.Y += enemyRec.Width
		for otherEnemy := range world.enemy {
			otherRec := util.CenterRectangle(world.position[otherEnemy], world.size[otherEnemy])
			otherRec.Width /= 2
			otherRec.Y += otherRec.Width
			if rl.CheckCollisionRecs(enemyRec, otherRec) {
				dir := util.Vector2Direction(world.position[enemy], world.position[otherEnemy])
				world.velocity[otherEnemy] = rl.Vector2Add(world.velocity[otherEnemy], rl.Vector2Scale(dir, 800*dt))
				world.velocity[enemy] = rl.Vector2Add(world.velocity[otherEnemy], rl.Vector2Scale(dir, -800*dt))
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
			if proj.typ == &projectileTypes.purpleBullet { // TODO: generalize this
				applyPoison(world, world.player, 3.5)
			}
			damageWithCooldown(world, world.player, proj.typ.damage, id)
			break
		}
	}
}
