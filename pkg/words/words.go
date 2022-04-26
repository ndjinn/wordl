package words

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Words struct {
	WordList    []string
	ReducedList []*string
}

func NewWordsFromFile(path string) *Words {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Words{WordList: words}
}

func (w Words) InWords(guess string) bool {
	upperGuess := strings.ToUpper(guess)

	for _, val := range w.WordList {
		if upperGuess == val {
			return true
		}
	}
	return false
}
