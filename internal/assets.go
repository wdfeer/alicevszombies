package internal

import rl "github.com/gen2brain/raylib-go/raylib"

var assets Assets

type Assets struct {
	Textures map[string]rl.Texture2D
}

func LoadAssets() {
	println("Starting to load assets...")
	assets = Assets{
		Textures: map[string]rl.Texture2D{
			"player": rl.LoadTexture("assets/player.png"),
		},
	}
	println("Assets loaded!")
}
