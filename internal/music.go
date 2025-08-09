package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var musicTracks = []string{"alice_stage", "alice_boss", "medicine", "nue", "kogasa", "tojiko"}

func updateMusic() {
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

	if playing != "" {
		rl.SetMusicVolume(assets.music[playing], options.MusicVolume)
		rl.UpdateMusicStream(assets.music[playing])
	} else {
		index := rand.Int() % len(musicTracks)
		rl.PlayMusicStream(assets.music[musicTracks[index]])
	}
}
