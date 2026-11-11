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

func (self *WorldGen) GetFoliageOnScreen(start rl.Vector2, end rl.Vector2) map[rl.Vector2]string {
	m := make(map[rl.Vector2]string, 0)
	// TODO: calculate foliage on screen with seed and biome data
	return m
}
