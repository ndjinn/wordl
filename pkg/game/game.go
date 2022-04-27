package game

import (
	"log"

	"github.com/ndjinn/wordl/pkg/words"
)

type GameConfig struct {
	WordFile    string
	MaxGuess    int
	Output      string
	Interactive bool
}

type GameState int

const (
	InProgress GameState = iota
	Victory
	Defeat
	Abandoned
)

type Game struct {
	config  GameConfig
	Words   words.Words
	target  string
	Guesses []Guess
	State   GameState
}

func (game *Game) MakeGuess(guess string) GameState {
	if !game.Words.InFullWords(guess) {
		log.Fatalln("The word '" + guess + "' is not a valid guess.")
	}

	newGuess := NewGuess(guess, game.target)
	game.Guesses = append(game.Guesses, *newGuess)

	if newGuess.Correct {
		return Victory
	}

	if game.config.MaxGuess <= len(game.Guesses) {
		return Defeat
	}

	return InProgress
}
