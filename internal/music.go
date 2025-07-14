package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var musicTracks = []string{"alice_stage", "alice_boss"}

func updateMusic() {
	played := ""

	for _, name := range musicTracks {
		if rl.IsMusicStreamPlaying(assets.music[name]) {
			played = name
			break
		}
	}

	if played != "" {
		rl.SetMusicVolume(assets.music[played], options.Volume)
		rl.UpdateMusicStream(assets.music[played])
	} else {
		index := rand.Int() % len(musicTracks)
		rl.PlayMusicStream(assets.music[musicTracks[index]])
	}
}
