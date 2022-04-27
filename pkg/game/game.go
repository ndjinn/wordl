package game

import (
	"github.com/ndjinn/wordl/pkg/words"
)

type GameConfig struct {
	WordFile    string
	MaxGuess    int
	Output      string
	Interactive bool
}

type Game struct {
	config GameConfig
	words  words.Words
	target string
	state  []Guess
}
