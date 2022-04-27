package game

import "fmt"

type Guess struct {
	Word   string
	Result []LetterResult
}

func NewGuess(guess string, target string) *Guess {
	g := &Guess{Word: guess}
	g.Result = LetterResultsFromDiff(guess, target)
	return g
}

func (g *Guess) printResultColor() {
	for _, v := range g.Result {
		v.printLetterResultColor()
	}
	fmt.Println()
}
