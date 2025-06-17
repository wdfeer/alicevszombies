package internal

type Difficulty = uint8

const (
	EASY    Difficulty = iota
	NORMAL  Difficulty = iota
	HARD    Difficulty = iota
	LUNATIC Difficulty = iota
)
