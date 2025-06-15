package util

// ModF divides the value by modder, returns the product of remainder times modder
func ModF(value float32, modder float32) float32 {
	divided := value / modder
	remainder := divided - float32(int64(divided))
	return remainder * modder
}
