package worldgen

import (
	"alicevszombies/internal/util"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WorldGen struct {
	seed uint32
}

func NewWorldGen() WorldGen {
	return WorldGen{
		rand.Uint32(),
	}
}

func (self *WorldGen) GetBiome(pos rl.Vector2) *Biome {
	return &biomes.plains
}

func (self *WorldGen) GetGrassTexture(pos rl.Vector2) string {
	return self.GetBiome(pos).grass
}

func (self *WorldGen) GetFoliageOnScreen(start rl.Vector2, end rl.Vector2) map[rl.Vector2]string {
	const SIZE = 32

	m := make(map[rl.Vector2]string, 0)

	start.X -= util.ModF(start.X, SIZE)
	start.Y -= util.ModF(start.Y, SIZE)
	end.X += util.ModF(end.X, SIZE)
	end.Y += util.ModF(end.Y, SIZE)

	for x := start.X; x < end.X; x += SIZE {
		for y := start.Y; y < end.Y; y += SIZE {
			pos := rl.Vector2{X: x, Y: y}
			seed := self.seed + util.Hash(uint32(pos.X)) + util.Hash(uint32(pos.Y))
			if util.Hash(seed)%100 < 8 {
				pos.X += float32(util.Hash(seed+1) % SIZE)
				pos.Y += float32(util.Hash(seed+2) % SIZE)

				biome := self.GetBiome(pos)
				foliage := biome.foliage[util.Hash(seed+3)%uint32(len(biome.foliage))]
				m[pos] = foliage
			}
		}
	}

	return m
}
