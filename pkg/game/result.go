package game

import (
	"encoding/json"
	"strings"

	"github.com/gookit/color"
)

type ResultColor color.Color

const (
	Right ResultColor = iota
	Misplaced
	Wrong
)

func (r *ResultColor) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*r = Wrong
	case "right":
		*r = Right
	case "misplaced":
		*r = Misplaced
	case "wrong":
		*r = Wrong
	}

	return nil
}

func (r ResultColor) MarshalJSON() ([]byte, error) {
	var s string
	switch r {
	default:
		s = "wrong"
	case Right:
		s = "right"
	case Misplaced:
		s = "misplaced"
	case Wrong:
		s = "wrong"
	}

	return json.Marshal(s)
}

func (r ResultColor) fgColor() color.Color {
	var c color.Color
	switch r {
	default:
		c = color.FgWhite
	case Right:
		c = color.FgWhite
	case Misplaced:
		c = color.FgBlack
	case Wrong:
		c = color.FgWhite
	}
	// Assuming all wrong colors are gray
	return c
}

func (r ResultColor) bgColor() color.Color {
	var c color.Color
	switch r {
	default:
		c = color.BgGray
	case Right:
		c = color.BgGreen
	case Misplaced:
		c = color.BgYellow
	case Wrong:
		c = color.BgGray
	}
	// Assuming all wrong colors are gray
	return c
}

func (r ResultColor) printResultColor(s string) {
	c := color.New(r.fgColor(), r.bgColor())
	c.Print(s)
}

type LetterResult struct {
	Letter string
	Result ResultColor
}

func (l *LetterResult) printLetterResultColor() {
	l.Result.printResultColor(l.Letter)
}

func LetterResultsFromDiff(guess string, target string) []LetterResult {
	diff := make([]LetterResult, len(guess))
	matched := make([]bool, len(target))

	for i := 0; i < len(guess); i++ {
		element := guess[i]
		result := Wrong
		if element == target[i] {
			result = Right
			matched[i] = true
		} else {
			for j := 0; j < len(target); j++ {
				if !matched[j] && element == target[j] {
					result = Misplaced
					matched[j] = true
					break
				}
			}
		}
		diff[i] = LetterResult{Letter: string(element), Result: result}
	}

	return diff
}
