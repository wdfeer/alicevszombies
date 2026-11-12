package util

// ModF divides the value by modder, returns the product of remainder times modder
func ModF(value float32, modder float32) float32 {
	divided := value / modder
	remainder := divided - float32(int64(divided))
	return remainder * modder
}

// Generate a pseudo-random number from seed
func Hash(seed uint32) uint32 {
	seed ^= seed >> 16
	seed *= 0x85ebca6b
	seed ^= seed >> 13
	seed *= 0xc2b2ae35
	seed ^= seed >> 16
	return seed
}
