package main

import (
	"fmt"

	"github.com/ndjinn/wordl/pkg/words"
)

func main() {
	words := words.NewWordsFromFile("./dict.txt")
	fmt.Println(len(words.FullList()))
	fmt.Println(words.FullList())
	fmt.Println(len(words.ReducedList()))
	fmt.Println(words.ReducedList())
}
