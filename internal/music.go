package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func updateMusic() {
	rl.SetSoundVolume(assets.sounds["alice_boss"], options.Volume)
	if !rl.IsSoundPlaying(assets.sounds["alice_boss"]) {
		rl.PlaySound(assets.sounds["alice_boss"])
	}
}
