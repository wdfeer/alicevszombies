package internal

import rl "github.com/gen2brain/raylib-go/raylib"

var assets Assets

type Assets struct {
	Textures map[string]rl.Texture2D
}

func LoadAssets() {
	println("Starting to load assets...")
	assets = Assets{
		Textures: make(map[string]rl.Texture2D),
	}

	loadTexture("player")
	loadTexture("player_walk0")
	loadTexture("player_walk1")
	loadTexture("cursor")
	loadTexture("doll1")

	println("Assets loaded!")
}

func loadTexture(name string) {
	assets.Textures[name] = rl.LoadTexture("assets/" + name + ".png")
}
