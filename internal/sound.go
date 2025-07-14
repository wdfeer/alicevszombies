package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func initMusic() {
	rl.PlaySound(assets.sounds["alice_boss"])
}

func updateMusic() {
	rl.SetSoundVolume(assets.sounds["alice_boss"], options.Volume)
	if !rl.IsSoundPlaying(assets.sounds["alice_boss"]) {
		rl.PlaySound(assets.sounds["alice_boss"])
	}
}

func playSound(name string) {
	rl.SetSoundVolume(assets.sounds[name], options.Volume)
	rl.PlaySound(assets.sounds[name])
}

func playSoundVolume(name string, volume float32) {
	if volume > 0 {
		rl.SetSoundVolume(assets.sounds[name], options.Volume*volume)
		rl.PlaySound(assets.sounds[name])
	}
}

func playSoundVolumePitch(name string, volume float32, pitch float32) {
	if volume > 0 {
		rl.SetSoundPitch(assets.sounds[name], pitch)
		playSoundVolume(name, volume)
	}
}
