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
