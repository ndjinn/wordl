package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ndjinn/wordl/pkg/words"
)

type GameConfig struct {
	WordFile    string
	MaxGuess    int
	Interactive bool
}

type GameState int

const (
	InProgress GameState = iota
	Victory
	Defeat
)

type Game struct {
	config  *GameConfig
	Words   *words.Words
	target  string
	Guesses []Guess
	State   GameState
}

func NewGame(conf *GameConfig) *Game {
	var guesses []Guess
	wordBank := words.NewWordsFromFile(conf.WordFile)

	return &Game{
		config:  conf,
		Words:   wordBank,
		target:  wordBank.RandomWord(),
		Guesses: guesses,
		State:   InProgress,
	}
}

func Play(cheat bool) {
	fmt.Println("Welcome to Word-L")
	config := &GameConfig{WordFile: "./dict.txt", MaxGuess: 6, Interactive: true}
	game := NewGame(config)

	if cheat {
		fmt.Printf("Cheater, your word is '%v'\n", game.target)
	}

	for game.State == InProgress {
		input := GetInputWord()
		valid := game.MakeGuess(input)

		game.PrintResults(valid)
	}
	game.GameStateMessage()
}

func (game *Game) MakeGuess(guess string) bool {
	if !game.Words.InFullWords(guess) {
		fmt.Printf("Word '%v' is not a valid guess. Please try again.\n", guess)
		return false
	}

	newGuess := NewGuess(guess, game.target)
	game.Guesses = append(game.Guesses, *newGuess)

	if newGuess.Correct {
		game.State = Victory
	}

	if game.config.MaxGuess <= len(game.Guesses) {
		game.State = Defeat
	}

	return true
}

func (game *Game) GetResultStrings() []string {
	var results []string
	for _, element := range game.Guesses {
		results = append(results, element.ResultColorString())
	}
	return results
}

func (game *Game) PrintResults(last bool) {
	var printStr string
	indent := "\t"
	results := game.GetResultStrings()
	if len(results) == 0 {
		return
	}

	fmt.Printf("Turn %v results:\n", len(results))
	if last {
		printStr = results[len(results)-1]
	} else {
		printStr = strings.Join(results, fmt.Sprintf("\n%v", indent))
	}
	fmt.Printf("%v%v\n", indent, printStr)
}

func (game *Game) GameStateMessage() {
	var msg string
	switch game.State {
	default:
		msg = "WORD-L"
	case InProgress:
		ml := game.config.MaxGuess - len(game.Guesses)
		msg = fmt.Sprintf("The game's not over yet!  You have %v turns left!", ml)
	case Victory:
		msg = "Congratulations!  You have won!"
	case Defeat:
		msg = fmt.Sprintf("Sorry, no moves left.  The word was '%v'", game.target)
	}
	fmt.Println(msg)
}

func GetInputWord() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return strings.ToUpper(text)
}
