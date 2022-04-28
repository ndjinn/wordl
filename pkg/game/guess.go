package game

import "fmt"

type Guess struct {
	Word    string
	Result  []LetterResult
	Correct bool
}

func NewGuess(guess string, target string) *Guess {
	g := &Guess{Word: guess}
	g.Result, g.Correct = LetterResultsFromDiff(guess, target)
	return g
}

func (g *Guess) ResultColorString() string {
	result := ""
	for _, v := range g.Result {
		result = result + v.letterResultColorString()
	}
	fmt.Println(result)
	return result
}
