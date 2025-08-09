package internal

import (
	"fmt"
	"time"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var assets = struct {
	renderTextures [2]rl.RenderTexture2D
	textures       map[string]rl.Texture2D
	breakdowns     map[string]TextureBreakdown
	sfx            map[string]rl.Sound
	music          map[string]rl.Music
	shaders        map[string]rl.Shader
}{
	textures:   make(map[string]rl.Texture2D),
	breakdowns: make(map[string]TextureBreakdown),
	sfx:        make(map[string]rl.Sound),
	music:      make(map[string]rl.Music),
	shaders:    make(map[string]rl.Shader),
}

func LoadAssets() {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("INFO: Loading assets took %s\n", elapsed)
	}()

	println("INFO: Starting to load assets...")

	rl.SetWindowIcon(*rl.LoadImage("assets/images/icon.png"))
	println("INFO: Icon loaded!")

	assets.renderTextures = [2]rl.RenderTexture2D{
		rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())),
		rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())),
	}
	println("INFO: Render Texture loaded!")

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
	loadTexture("chunky_zombie")
	loadTexture("chunky_zombie_walk0")
	loadTexture("chunky_zombie_walk1")
	loadTexture("purple_zombie")
	loadTexture("purple_zombie_walk0")
	loadTexture("purple_zombie_walk1")
	loadTexture("blue_zombie")
	loadTexture("blue_zombie_walk0")
	loadTexture("blue_zombie_walk1")
	loadTexture("nerium_girl")
	loadTexture("nerium_girl_walk0")
	loadTexture("nerium_girl_walk1")
	loadTexture("zombie_fairy")
	loadTexture("medicine")
	loadTexture("medicine_walk0")
	loadTexture("medicine_walk1")
	loadTextureAndFlipped("kogasa")
	loadTextureAndFlipped("kogasa_walk0")
	loadTextureAndFlipped("kogasa_walk1")
	loadTextureAndFlipped("tojiko")
	loadTextureAndFlipped("nue")
	loadTextureAndFlipped("nue_walk0")
	loadTextureAndFlipped("nue_walk1")
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
	loadTexture("lightning_bolt")

	println("INFO: Textures loaded!")

	loadBreakdown("player")
	loadBreakdown("zombie")
	loadBreakdown("small_zombie")
	loadBreakdown("chunky_zombie")
	loadBreakdown("purple_zombie")
	loadBreakdown("blue_zombie")
	loadBreakdown("nerium_girl")
	loadBreakdown("zombie_fairy")
	loadBreakdown("medicine")
	loadBreakdown("kogasa")
	loadBreakdown("tojiko")
	loadBreakdown("nue")

	loadBreakdown("doll_sword")
	loadBreakdown("doll_lance")
	loadBreakdown("doll_scythe")
	loadBreakdown("doll_destruction")
	loadBreakdown("doll_knife")
	loadBreakdown("doll_magician")

	println("INFO: Texture Breakdowns loaded!")

	assets.shaders["bloom"] = rl.LoadShader("", "assets/shaders/bloom.fs")
	assets.shaders["chroma_abberation"] = rl.LoadShader("", "assets/shaders/chroma_abberation.fs")
	println("INFO: Shaders loaded!")

	rl.InitAudioDevice()
	loadSFX("player_hit")
	loadSFX("enemy_hit")
	loadSFX("enemy_kill")
	loadSFX("boss_spawn")
	println("INFO: Sounds loaded!")

	for _, name := range musicTracks {
		loadMusic(name)
		println("INFO: Music,", name, " loaded!")
	}
	println("INFO: Sounds loaded!")

	raygui.LoadStyle("assets/style.rgs")
	updateUIScale()
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT, raygui.TEXT_ALIGN_CENTER)
	println("INFO: Raygui style loaded!")
}

func UnloadAssets() {
	for _, texture := range assets.textures {
		rl.UnloadTexture(texture)
	}

	rl.CloseAudioDevice()
	for _, sound := range assets.sfx {
		rl.UnloadSound(sound)
	}
	for _, music := range assets.music {
		rl.UnloadMusicStream(music)
	}

	for _, shader := range assets.shaders {
		rl.UnloadShader(shader)
	}

	for _, v := range assets.renderTextures {
		rl.UnloadRenderTexture(v)
	}
}

func loadTexture(name string) {
	assets.textures[name] = rl.LoadTexture("assets/images/" + name + ".png")
}

const FlippedSuffix = "_fliph"

func loadTextureAndFlipped(name string) {
	image := rl.LoadImage("assets/images/" + name + ".png")
	assets.textures[name] = rl.LoadTextureFromImage(image)
	rl.ImageFlipHorizontal(image)
	assets.textures[name+FlippedSuffix] = rl.LoadTextureFromImage(image)
}

func loadSFX(name string) {
	assets.sfx[name] = rl.LoadSound("assets/sfx/" + name + ".wav")
}

func loadMusic(name string) {
	assets.music[name] = rl.LoadMusicStream("assets/music/" + name + ".ogg")
}
