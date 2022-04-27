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
	target := "great"

	guess1 := "goats"
	expected1 := `[
		{"Letter": "g", "Result": "right"},
		{"Letter": "o", "Result": "wrong"},
		{"Letter": "a", "Result": "misplaced"},
		{"Letter": "t", "Result": "misplaced"},
		{"Letter": "s", "Result": "wrong"}
	]`
	result1, correct1 := LetterResultsFromDiff(guess1, target)
	json1, _ := json.Marshal(result1)
	assert.JSONEq(t, expected1, string(json1))
	assert.False(t, correct1)

	guess2 := "greet"
	expected2 := `[
		{"Letter": "g", "Result": "right"},
		{"Letter": "r", "Result": "right"},
		{"Letter": "e", "Result": "right"},
		{"Letter": "e", "Result": "wrong"},
		{"Letter": "t", "Result": "right"}
	]`
	result2, correct2 := LetterResultsFromDiff(guess2, target)
	json2, _ := json.Marshal(result2)
	assert.JSONEq(t, expected2, string(json2))
	assert.False(t, correct2)

	guess3 := "great"
	expected3 := `[
		{"Letter": "g", "Result": "right"},
		{"Letter": "r", "Result": "right"},
		{"Letter": "e", "Result": "right"},
		{"Letter": "a", "Result": "right"},
		{"Letter": "t", "Result": "right"}
	]`
	result3, correct3 := LetterResultsFromDiff(guess3, target)
	json3, _ := json.Marshal(result3)
	assert.JSONEq(t, expected3, string(json3))
	assert.True(t, correct3)
}
