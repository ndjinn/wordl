package game

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func TestFgColor(t *testing.T) {
	assert.Equal(t, Right.fgColor(), color.FgWhite)
	assert.Equal(t, Misplaced.fgColor(), color.FgBlack)
	assert.Equal(t, Wrong.fgColor(), color.FgWhite)
}

func TestBgColor(t *testing.T) {
	assert.Equal(t, Right.bgColor(), color.BgGreen)
	assert.Equal(t, Misplaced.bgColor(), color.BgYellow)
	assert.Equal(t, Wrong.bgColor(), color.BgGray)
}

func TestResultColorMarshalJSON(t *testing.T) {
	r, _ := json.Marshal(Right)
	m, _ := json.Marshal(Misplaced)
	w, _ := json.Marshal(Wrong)

	assert.JSONEq(t, `"right"`, string(r))
	assert.JSONEq(t, `"misplaced"`, string(m))
	assert.JSONEq(t, `"wrong"`, string(w))
}

func TestResultColorUnmarshalJSON(t *testing.T) {
	rjs := `"right"`
	mjs := `"misplaced"`
	wjs := `"wrong"`

	var r ResultColor
	var m ResultColor
	var w ResultColor

	if err := json.Unmarshal([]byte(rjs), &r); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(mjs), &m); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(wjs), &w); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, Right, r)
	assert.Equal(t, Misplaced, m)
	assert.Equal(t, Wrong, w)
}

func TestLetterResultsFromDiff(t *testing.T) {
	target := "GREAT"

	// Simple match
	guess1 := "GOATS"
	expected1 := `[
		{"Letter": "G", "Result": "right"},
		{"Letter": "O", "Result": "wrong"},
		{"Letter": "A", "Result": "misplaced"},
		{"Letter": "T", "Result": "misplaced"},
		{"Letter": "S", "Result": "wrong"}
	]`
	result1, correct1 := LetterResultsFromDiff(guess1, target)
	json1, _ := json.Marshal(result1)
	assert.JSONEq(t, expected1, string(json1))
	assert.False(t, correct1)

	// Testing double letter
	guess2 := "GREET"
	expected2 := `[
		{"Letter": "G", "Result": "right"},
		{"Letter": "R", "Result": "right"},
		{"Letter": "E", "Result": "right"},
		{"Letter": "E", "Result": "wrong"},
		{"Letter": "T", "Result": "right"}
	]`
	result2, correct2 := LetterResultsFromDiff(guess2, target)
	json2, _ := json.Marshal(result2)
	assert.JSONEq(t, expected2, string(json2))
	assert.False(t, correct2)

	// Testing correct guess
	guess3 := "GREAT"
	expected3 := `[
		{"Letter": "G", "Result": "right"},
		{"Letter": "R", "Result": "right"},
		{"Letter": "E", "Result": "right"},
		{"Letter": "A", "Result": "right"},
		{"Letter": "T", "Result": "right"}
	]`
	result3, correct3 := LetterResultsFromDiff(guess3, target)
	json3, _ := json.Marshal(result3)
	assert.JSONEq(t, expected3, string(json3))
	assert.True(t, correct3)

	//Test double letters where the first is wrong and the second is correct
	target4 := "BREWS"
	guess4 := "BEERS"
	expected4 := `[
		{"Letter": "B", "Result": "right"},
		{"Letter": "E", "Result": "wrong"},
		{"Letter": "E", "Result": "right"},
		{"Letter": "R", "Result": "misplaced"},
		{"Letter": "S", "Result": "right"}
	]`
	result4, correct4 := LetterResultsFromDiff(guess4, target4)
	json4, _ := json.Marshal(result4)
	assert.JSONEq(t, expected4, string(json4))
	assert.False(t, correct4)

	//Test double letters where the first is misplaced and the second is correct
	target5 := "BLEED"
	guess5 := "BEERS"
	expected5 := `[
		{"Letter": "B", "Result": "right"},
		{"Letter": "E", "Result": "misplaced"},
		{"Letter": "E", "Result": "right"},
		{"Letter": "R", "Result": "wrong"},
		{"Letter": "S", "Result": "wrong"}
	]`
	result5, correct5 := LetterResultsFromDiff(guess5, target5)
	json5, _ := json.Marshal(result5)
	assert.JSONEq(t, expected5, string(json5))
	assert.False(t, correct5)
}
