package internal

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var assets = struct {
	textures     map[string]rl.Texture2D
	deathEffects map[string]DeathEffectAsset
	sounds       map[string]rl.Sound
}{
	textures:     make(map[string]rl.Texture2D),
	deathEffects: make(map[string]DeathEffectAsset),
	sounds:       make(map[string]rl.Sound),
}

func LoadAssets() {
	println("INFO: Starting to load assets...")

	rl.SetWindowIcon(*rl.LoadImage("assets/icon.png"))
	println("INFO: Icon loaded!")

	loadTexture("player")
	loadTexture("player_walk0")
	loadTexture("player_walk1")
	loadTexture("cursor0")
	loadTexture("cursor1")

	loadTextureAndFlipped("doll_sword")
	loadTextureAndFlipped("doll_lance")
	loadTextureAndFlipped("doll_scythe")
	loadTextureAndFlipped("doll_destruction")
	loadTexture("doll_knife")
	loadTexture("doll_magician")

	loadTexture("zombie")
	loadTexture("zombie_walk0")
	loadTexture("zombie_walk1")
	loadTexture("small_zombie")
	loadTexture("small_zombie_walk0")
	loadTexture("small_zombie_walk1")
	loadTexture("purple_zombie")
	loadTexture("purple_zombie_walk0")
	loadTexture("purple_zombie_walk1")
	loadTexture("blue_zombie")
	loadTexture("blue_zombie_walk0")
	loadTexture("blue_zombie_walk1")
	loadTexture("medicine")
	loadTexture("medicine_walk0")
	loadTexture("medicine_walk1")
	loadTextureAndFlipped("kogasa")
	loadTextureAndFlipped("kogasa_walk0")
	loadTextureAndFlipped("kogasa_walk1")
	loadTexture("heal_icon")
	loadTexture("doll_icon")
	loadTexture("pitem_icon")
	loadTexture("unique_upgrade_icon")
	loadTexture("grass")

	loadTexture("knife")
	loadTexture("red_bullet")
	loadTexture("magic_missile")
	loadTexture("purple_bullet")
	loadTexture("blue_bullet")

	println("INFO: Textures loaded!")

	loadDeathEffect("player")
	loadDeathEffect("zombie")
	loadDeathEffect("small_zombie")
	loadDeathEffect("purple_zombie")
	loadDeathEffect("blue_zombie")
	loadDeathEffect("medicine")
	loadDeathEffect("kogasa")

	println("INFO: Death Effects loaded!")

	rl.InitAudioDevice()
	loadSound("player_hit")
	loadSound("enemy_hit")
	println("INFO: Sounds loaded!")

	raygui.LoadStyle("assets/style.rgs")
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, 80)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, 8)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, raygui.TEXT_ALIGN_CENTER)
	println("INFO: Raygui style loaded!")
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

const FlippedSuffix = "_fliph"

func loadTextureAndFlipped(name string) {
	image := rl.LoadImage("assets/" + name + ".png")
	assets.textures[name] = rl.LoadTextureFromImage(image)
	rl.ImageFlipHorizontal(image)
	assets.textures[name+FlippedSuffix] = rl.LoadTextureFromImage(image)
}

func loadSound(name string) {
	assets.sounds[name] = rl.LoadSound("assets/" + name + ".wav")
}
