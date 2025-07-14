package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateMusic() {
	rl.SetMusicVolume(assets.music["alice_boss"], options.Volume)
	if !rl.IsMusicStreamPlaying(assets.music["alice_boss"]) {
		rl.PlayMusicStream(assets.music["alice_boss"])
	}

	rl.UpdateMusicStream(assets.music["alice_boss"])
}
