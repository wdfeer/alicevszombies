package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var musicTracks = []string{"alice_stage", "alice_boss", "medicine", "nue", "kogasa", "tojiko"}

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
		if world.enemySpawner.wave < 10 {
			musicToPlay = "alice_stage"
		} else {
			musicToPlay = "alice_boss"
		}
	}

	playing := ""

	for _, name := range musicTracks {
		if rl.IsMusicStreamPlaying(assets.music[name]) {
			playing = name
			break
		}
	}

	if playing != musicToPlay {
		if playing != "" {
			rl.StopMusicStream(assets.music[playing])
		}
		rl.PlayMusicStream(assets.music[musicToPlay])
		playing = musicToPlay
	}

	rl.SetMusicVolume(assets.music[playing], options.MusicVolume)
	rl.UpdateMusicStream(assets.music[playing])
}
