package worldgen

import rl "github.com/gen2brain/raylib-go/raylib"

type WorldGen struct {
	seed int
}

func (self *WorldGen) GetGrassTexture(pos rl.Vector2) string {
	return "grass"
}
