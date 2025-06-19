package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func playSound(name string) {
	rl.SetSoundVolume(assets.sounds[name], soundVolume)
	rl.PlaySound(assets.sounds[name])
}

func playSoundVolume(name string, volume float32) {
	if volume > 0 {
		rl.SetSoundVolume(assets.sounds[name], soundVolume*volume)
		rl.PlaySound(assets.sounds[name])
	}
}

func playSoundVolumePitch(name string, volume float32, pitch float32) {
	if volume > 0 {
		rl.SetSoundPitch(assets.sounds[name], pitch)
		playSoundVolume(name, volume)
	}
}
