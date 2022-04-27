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

func (g *Guess) ResultColorString() string {
	result := ""
	for _, v := range g.Result {
		result = result + v.letterResultColorString()
	}
	fmt.Println(result)
	return result
}
