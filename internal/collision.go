package internal

import (
	"alicevszombies/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateCollisions(world *World) {
	playerPos := world.position[world.player]
	playerRec := util.CenterRectangle(playerPos, world.size[world.player])

	for enemy, enemyType := range world.enemy {
		enemyPos := world.position[enemy]
		enemyRec := util.CenterRectangle(enemyPos, world.size[enemy])

		// Enemy -> Player
		if util.CheckCollisionRecs(playerRec, enemyRec) {
			if world.difficulty == LUNATIC {
				applyStatus(world, world.player, Bleed, 5.5)
			}
			damageWithCooldown(world, world.player, 1, enemy)

			dir := util.Vector2Direction(playerPos, enemyPos)
			world.velocity[enemy] = rl.Vector2Add(world.velocity[enemy], rl.Vector2Scale(dir, 800*dt))
		}

		// Doll -> Enemy
		for doll, typ := range world.doll {
			if typ.contactDamage <= 0 || typ.size.X <= 0 {
				continue
			}

			if util.CheckCollisionCenteredVsRec(world.position[doll], world.size[doll], enemyRec) {
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

			if util.CheckCollisionCenteredVsRec(world.position[id], proj.typ.size, enemyRec) {
				dmg := proj.typ.damage + (float32(world.playerData.upgrades[&DollDamage]/2) / 4)
				if proj.typ.deleteOnHit {
					damage(world, enemy, dmg)
					world.deleteEntity(id)
				} else {
					damageWithCooldown(world, enemy, dmg, id)
				}
				break
			}
		}

		// Enemy <-> Enemy, knockback only
		if !enemyType.flying {
			enemyRec.Width /= 2
			enemyRec.Y += enemyRec.Width
			for otherEnemy, otherType := range world.enemy {
				if otherType.flying {
					continue
				}

				otherPos := world.position[otherEnemy]
				otherRec := util.CenterRectangle(otherPos, world.size[otherEnemy])
				otherRec.Width /= 2
				otherRec.Y += otherRec.Width
				if util.CheckCollisionRecs(enemyRec, otherRec) {
					dir := util.Vector2Direction(enemyPos, otherPos)
					world.velocity[otherEnemy] = rl.Vector2Add(world.velocity[otherEnemy], rl.Vector2Scale(dir, 800*dt))
					world.velocity[enemy] = rl.Vector2Add(world.velocity[otherEnemy], rl.Vector2Scale(dir, -800*dt))
				}
			}
		}
	}

	statusDuration := float32(3.5)
	if world.difficulty > NORMAL {
		statusDuration = 6.5
	}

	// Projectile -> Player
	for id, proj := range world.projectile {
		if !proj.typ.hostile {
			continue
		}

		projRec := util.CenterRectangle(world.position[id], proj.typ.size)
		if util.CheckCollisionRecs(playerRec, projRec) {
			if proj.typ == &projectileTypes.purpleBullet { // TODO: generalize this
				applyStatus(world, world.player, Poison, statusDuration)
			} else if proj.typ == &projectileTypes.blueBullet {
				applyStatus(world, world.player, Slow, statusDuration)
			}

			damageWithCooldown(world, world.player, proj.typ.damage, id)
			break
		}

		// Projectile -> Doll
		if proj.typ == &projectileTypes.blueBullet {
			for doll := range world.doll {
				dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
				if util.CheckCollisionRecs(projRec, dollRec) {
					applyStatus(world, doll, Slow, statusDuration)
				}
			}
		}
	}
}
