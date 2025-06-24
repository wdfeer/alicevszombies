package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	typ      *ProjectileType
	timeLeft float32
}

func newProjectile(world *World, pos rl.Vector2, vel rl.Vector2, typ *ProjectileType) Entity {
	id := world.newEntity()
	world.position[id] = pos
	world.velocity[id] = vel
	world.projectile[id] = Projectile{
		typ:      typ,
		timeLeft: 4,
	}
	world.texture[id] = typ.texture
	world.size[id] = typ.size
	return id
}

func updateProjectiles(world *World) {
	for id, proj := range world.projectile {
		timeLeft := proj.timeLeft - dt
		if timeLeft <= 0 {
			world.deleteEntity(id)
		} else {
			world.projectile[id] = Projectile{proj.typ, proj.timeLeft - dt}
		}
	}
}

type ProjectileType struct {
	hostile     bool
	damage      float32
	texture     string
	size        rl.Vector2
	deleteOnHit bool
}

var projectileTypes = struct {
	knife        ProjectileType
	purpleBullet ProjectileType
	magicMissile ProjectileType
	redBullet    ProjectileType
}{
	knife: ProjectileType{
		hostile:     false,
		damage:      1,
		texture:     "knife",
		size:        rl.Vector2{X: 4, Y: 4},
		deleteOnHit: true,
	},
	purpleBullet: ProjectileType{
		hostile: true,
		damage:  5,
		texture: "purple_bullet",
		size:    rl.Vector2{X: 4, Y: 4},
	},
	magicMissile: ProjectileType{
		hostile:     false,
		damage:      2,
		texture:     "magic_missile",
		size:        rl.Vector2{X: 4, Y: 4},
		deleteOnHit: false,
	},
	redBullet: ProjectileType{
		hostile: false,
		damage:  5,
		texture: "red_bullet",
		size:    rl.Vector2{X: 4, Y: 4},
	},
}
