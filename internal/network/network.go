package network

type NetMode = int8

const (
	Singleplayer NetMode = iota
	Multiplayer
)

var Mode NetMode = Singleplayer
