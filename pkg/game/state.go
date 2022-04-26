package game

import "github.com/ndjinn/wordl/pkg/words"

type GameConfig struct {
	wordFile string
	maxGuess uint
}

type Game struct {
	config GameConfig
	words  words.Words
	target string
	state  []Guess
}
