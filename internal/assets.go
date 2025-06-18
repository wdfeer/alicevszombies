package internal

import rl "github.com/gen2brain/raylib-go/raylib"

var assets Assets

type Assets struct {
	textures     map[string]rl.Texture2D
	deathEffects map[string]DeathEffectAsset
	sounds       map[string]rl.Sound
}

func LoadAssets() {

	println("Starting to load assets...")
	assets = Assets{
		textures:     make(map[string]rl.Texture2D),
		deathEffects: make(map[string]DeathEffectAsset),
		sounds:       make(map[string]rl.Sound),
	}

	loadTexture("player")
	loadTexture("player_walk0")
	loadTexture("player_walk1")
	loadTexture("cursor")

	loadTextureAndFlipped("doll_sword")
	loadTextureAndFlipped("doll_lance")
	loadTexture("doll_knife")

	loadTexture("zombie")
	loadTexture("zombie_walk0")
	loadTexture("zombie_walk1")
	loadTexture("red_zombie")
	loadTexture("red_zombie_walk0")
	loadTexture("red_zombie_walk1")
	loadTexture("medicine")
	loadTexture("medicine_walk0")
	loadTexture("medicine_walk1")
	loadTexture("heal_icon")
	loadTexture("doll_icon")
	loadTexture("pitem_icon")
	loadTexture("grass")

	loadTexture("knife")
	loadTexture("red_bullet")

	println("INFO: Textures loaded!")

	loadDeathEffect("player")
	loadDeathEffect("zombie")
	loadDeathEffect("red_zombie")
	loadDeathEffect("medicine")

	println("INFO: Death Effects loaded!")

	rl.InitAudioDevice()
	loadSound("player_hit")
	println("INFO: Sounds loaded!")
}

func UnloadAssets() {
	for _, texture := range assets.textures {
		rl.UnloadTexture(texture)
	}

	rl.CloseAudioDevice()
	for _, sound := range assets.sounds {
		rl.UnloadSound(sound)
	}
}

func loadTexture(name string) {
	assets.textures[name] = rl.LoadTexture("assets/" + name + ".png")
}

func loadTextureAndFlipped(name string) {
	image := rl.LoadImage("assets/" + name + ".png")
	assets.textures[name] = rl.LoadTextureFromImage(image)
	rl.ImageFlipHorizontal(image)
	assets.textures[name+"_fliph"] = rl.LoadTextureFromImage(image)
}

func loadSound(name string) {
	assets.sounds[name] = rl.LoadSound("assets/" + name + ".wav")
}
