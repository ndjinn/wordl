package main

import (
	"os"

	"github.com/ndjinn/wordl/pkg/game"
)

func main() {
	cheat := false
	if len(os.Args) > 2 && os.Args[1] == "cheat" {
		cheat = true
	}
	game.Play(cheat)
}
