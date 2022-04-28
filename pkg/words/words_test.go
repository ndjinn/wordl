package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lowercaseFile string = "./testfiles/lowercase.txt"
var randomcaseFile string = "./testfiles/randomcase.txt"
var uppercaseFile string = "./testfiles/uppercase.txt"

func TestWordsFromFile(t *testing.T) {
	words := []string{"FOO", "BAR", "BAZ", "QUX"}

	lowercase := NewWordsFromFile(lowercaseFile)
	randomcase := NewWordsFromFile(randomcaseFile)
	uppercase := NewWordsFromFile(uppercaseFile)

	assert.ElementsMatch(t, words, lowercase.FullList())
	assert.ElementsMatch(t, words, randomcase.FullList())
	assert.ElementsMatch(t, words, uppercase.FullList())
	assert.ElementsMatch(t, uppercase.ReducedList(), uppercase.FullList())
}

func TestInWords(t *testing.T) {
	expectedTrue := []string{"FOO", "foo", "FoO"}
	expectedFalse := []string{"CAT", "cat", "CaT", ""}

	words := NewWordsFromFile(lowercaseFile)

	for _, val := range expectedTrue {
		assert.True(t, words.InFullWords(val))
	}

	for _, val := range expectedFalse {
		assert.False(t, words.InFullWords(val))
	}
}
