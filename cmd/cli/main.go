package main

import (
	"fmt"

	"github.com/ndjinn/wordl/pkg/words"
)

func main() {
	words := words.NewWordsFromFile("./dict.txt")
	fmt.Println(words.WordList)
	fmt.Println(len(words.WordList))
	fmt.Println(words.ReducedList)
}
