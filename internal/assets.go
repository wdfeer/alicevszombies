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
	loadTextureAndFlipped("doll1")
	loadTexture("zombie")
	loadTexture("zombie_walk0")
	loadTexture("zombie_walk1")

	loadTexture("heal_icon")
	loadTexture("doll_icon")
	loadTexture("pitem_icon")

	println("Assets loaded!")
}

func loadTexture(name string) {
	assets.Textures[name] = rl.LoadTexture("assets/" + name + ".png")
}

func loadTextureAndFlipped(name string) {
	image := rl.LoadImage("assets/" + name + ".png")
	assets.Textures[name] = rl.LoadTextureFromImage(image)
	rl.ImageFlipHorizontal(image)
	assets.Textures[name+"_fliph"] = rl.LoadTextureFromImage(image)
}
