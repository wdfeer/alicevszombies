package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var musicTracks = map[string]float32{
	"alice_stage": 0,
	"alice_boss":  0,
	"medicine":    0,
	"nue":         0,
	"kogasa":      0,
	"tojiko":      0,
}

func updateMusic(world *World) {
	if options.MusicVolume == 0 {
		return
	}

	var musicToPlay string
	for _, v := range world.enemy {
		if v.spawnData.boss {
			switch v {
			case &enemyTypes.medicine:
				musicToPlay = "medicine"
			case &enemyTypes.nue:
				musicToPlay = "nue"
			case &enemyTypes.kogasa:
				musicToPlay = "kogasa"
			case &enemyTypes.tojiko:
				musicToPlay = "tojiko"
			}
			break
		}
	}
	if musicToPlay == "" {
		if world.enemySpawner.wave < 11 {
			musicToPlay = "alice_stage"
		} else {
			musicToPlay = "alice_boss"
		}
	}

	for name, vol := range musicTracks {
		if rl.IsMusicStreamPlaying(assets.music[name]) {
			musicTracks[name] = rl.Clamp(vol, 0, 1)
			if vol == 0 {
				rl.StopMusicStream(assets.music[name])
			} else {
				rl.SetMusicVolume(assets.music[name], options.MusicVolume*vol)
				rl.UpdateMusicStream(assets.music[name])
				musicTracks[name] -= dt
			}
		} else if vol > 0 {
			rl.PlayMusicStream(assets.music[name])
		}
	}

	musicTracks[musicToPlay] += dt * 1.5
}
