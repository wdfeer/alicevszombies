package worldgen

import rl "github.com/gen2brain/raylib-go/raylib"

type WorldGen struct {
	seed int
}

func (self *WorldGen) GetBiome(pos rl.Vector2) *Biome {
	return &biomes.plains
}

func (self *WorldGen) GetGrassTexture(pos rl.Vector2) string {
	return self.GetBiome(pos).grass
}
