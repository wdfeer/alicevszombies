package internal

import (
	"alicevszombies/internal/util"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collision struct {
	attacker Entity
	victim   Entity
}

type CollisionType = uint8

const (
	EnemyPlayer CollisionType = iota
	DollEnemy
	ProjectileEnemy
	EnemyEnemy
	ProjectilePlayer
	ProjectileDoll
)

func checkCollisions(world *World, typ CollisionType) []Collision {
	var collisions []Collision
	playerRec := util.CenterRectangle(world.position[world.player], world.size[world.player])

	switch typ {
	case EnemyPlayer:
		for enemy := range world.enemy {
			enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])
			if rl.CheckCollisionRecs(playerRec, enemyRec) {
				collisions = append(collisions, Collision{enemy, world.player})
			}
		}

	case DollEnemy:
		for enemy := range world.enemy {
			enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])
			for doll, dollType := range world.doll {
				if dollType.contactDamage <= 0 || dollType.size.X <= 0 {
					continue
				}
				dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
				if rl.CheckCollisionRecs(dollRec, enemyRec) {
					collisions = append(collisions, Collision{doll, enemy})
					break
				}
			}
		}

	case ProjectileEnemy:
		for enemy := range world.enemy {
			enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])
			for id, proj := range world.projectile {
				if proj.typ.hostile {
					continue
				}
				projRec := util.CenterRectangle(world.position[id], proj.typ.size)
				if rl.CheckCollisionRecs(enemyRec, projRec) {
					collisions = append(collisions, Collision{id, enemy})
					break
				}
			}
		}

	case EnemyEnemy:
		for enemy := range world.enemy {
			enemyRec := util.CenterRectangle(world.position[enemy], world.size[enemy])
			enemyRec.Width /= 2
			enemyRec.Y += enemyRec.Width
			for otherEnemy := range world.enemy {
				if otherEnemy == enemy {
					continue
				}
				otherRec := util.CenterRectangle(world.position[otherEnemy], world.size[otherEnemy])
				otherRec.Width /= 2
				otherRec.Y += otherRec.Width
				if rl.CheckCollisionRecs(enemyRec, otherRec) {
					collisions = append(collisions, Collision{enemy, otherEnemy})
				}
			}
		}

	case ProjectilePlayer, ProjectileDoll:
		for id, proj := range world.projectile {
			if !proj.typ.hostile {
				continue
			}
			projRec := util.CenterRectangle(world.position[id], proj.typ.size)

			if typ == ProjectilePlayer {
				if rl.CheckCollisionRecs(playerRec, projRec) {
					collisions = append(collisions, Collision{id, world.player})
				}
			}

			if typ == ProjectileDoll && proj.typ == &projectileTypes.blueBullet {
				for doll := range world.doll {
					dollRec := util.CenterRectangle(world.position[doll], world.size[doll])
					if rl.CheckCollisionRecs(projRec, dollRec) {
						collisions = append(collisions, Collision{id, doll})
					}
				}
			}
		}
	}

	return collisions
}

func processCollisions(world *World, collisions []Collision, typ CollisionType) {
	statusDuration := float32(3.5)
	if world.difficulty > NORMAL {
		statusDuration = 6.5
	}

	switch typ {
	case EnemyPlayer:
		for _, c := range collisions {
			damageWithCooldown(world, c.victim, 1, c.attacker)
			dir := util.Vector2Direction(world.position[c.victim], world.position[c.attacker])
			world.velocity[c.attacker] = rl.Vector2Add(world.velocity[c.attacker], rl.Vector2Scale(dir, 800*dt))
		}

	case DollEnemy:
		for _, c := range collisions {
			dollType := world.doll[c.attacker]
			dmg := dollType.contactDamage + float32(world.playerData.upgrades[&DollDamage])/4
			if dollType == &dollTypes.scytheDoll {
				dmg += float32(world.playerData.upgrades[&DollDamage]) / 4
			}
			damageWithCooldown(world, c.victim, dmg, c.attacker)
		}

	case ProjectileEnemy:
		for _, c := range collisions {
			proj := world.projectile[c.attacker]
			dmg := proj.typ.damage + float32(world.playerData.upgrades[&DollDamage])/8
			if proj.typ.deleteOnHit {
				damage(world, c.victim, dmg)
				world.deleteEntity(c.attacker)
			} else {
				damageWithCooldown(world, c.victim, dmg, c.attacker)
			}
		}

	case EnemyEnemy:
		for _, c := range collisions {
			dir := util.Vector2Direction(world.position[c.attacker], world.position[c.victim])
			world.velocity[c.victim] = rl.Vector2Add(world.velocity[c.victim], rl.Vector2Scale(dir, 800*dt))
			world.velocity[c.attacker] = rl.Vector2Add(world.velocity[c.victim], rl.Vector2Scale(dir, -800*dt))
		}

	case ProjectilePlayer:
		for _, c := range collisions {
			proj := world.projectile[c.attacker]
			if proj.typ == &projectileTypes.purpleBullet {
				applyPoison(world, c.victim, statusDuration)
			} else if proj.typ == &projectileTypes.blueBullet {
				applySlow(world, c.victim, statusDuration)
			}
			damageWithCooldown(world, c.victim, proj.typ.damage, c.attacker)
		}

	case ProjectileDoll:
		for _, c := range collisions {
			if world.projectile[c.attacker].typ == &projectileTypes.blueBullet {
				applySlow(world, c.victim, statusDuration)
			}
		}
	}
}

func updateCollisions(world *World) {
	types := []CollisionType{
		EnemyPlayer,
		DollEnemy,
		ProjectileEnemy,
		EnemyEnemy,
		ProjectilePlayer,
		ProjectileDoll,
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	collisionMap := make(map[CollisionType][]Collision)

	for _, t := range types {
		wg.Add(1)
		go func() {
			defer wg.Done()

			collisions := checkCollisions(world, t)
			mu.Lock()
			collisionMap[t] = collisions
			mu.Unlock()
		}()
	}

	wg.Wait()

	for _, t := range types {
		processCollisions(world, collisionMap[t], t)
	}
}
