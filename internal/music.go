package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var musicTracks = []string{"alice_stage", "alice_boss", "medicine", "nue", "kogasa", "tojiko"}

func updateMusic(world *World) {
	if options.MusicVolume == 0 {
		return
	}

	playing := ""

	for _, name := range musicTracks {
		if rl.IsMusicStreamPlaying(assets.music[name]) {
			playing = name
			break
		}
	}

	for _, v := range world.enemy {
		if v.spawnData.boss {
			music := v.texture
			if playing != music {
				rl.PlayMusicStream(assets.music[music])
			}
			break
		}
	}

	if playing != "" {
		rl.SetMusicVolume(assets.music[playing], options.MusicVolume)
		rl.UpdateMusicStream(assets.music[playing])
	} else {
		index := rand.Int() % 2 // either alice_stage or alice_boss
		rl.PlayMusicStream(assets.music[musicTracks[index]])
	}
}
