package internal

type ShootPattern struct {
	projectile *ProjectileType
	cooldown   float32
	typ        ShootType

	// Only used for some shoot types
	count uint8
}

type ShootType = uint8

const (
	Direct ShootType = iota
	Circle
	Spread
)
