package worldgen

var biomes = struct {
	plains Biome
}{
	plains: Biome{
		grass: "green_grass",
		foliage: []string{
			"green_foliage0",
			"green_foliage1",
		},
	},
}
