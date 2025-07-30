package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func playSound(name string) {
	rl.SetSoundVolume(assets.sfx[name], options.SoundVolume)
	rl.PlaySound(assets.sfx[name])
}

func playSoundVolume(name string, volume float32) {
	if volume > 0 {
		rl.SetSoundVolume(assets.sfx[name], options.SoundVolume*volume)
		rl.PlaySound(assets.sfx[name])
	}
}

func playSoundVolumePitch(name string, volume float32, pitch float32) {
	if volume > 0 {
		rl.SetSoundPitch(assets.sfx[name], pitch)
		playSoundVolume(name, volume)
	}
}
