package words

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ndjinn/wordl/pkg/common"
)

type Words struct {
	fullList    []string
	reducedList []*string
}

func NewWords(wl []string) *Words {
	var rl []*string

	for i := range wl {
		rl = append(rl, &wl[i])
	}

	return &Words{fullList: wl, reducedList: rl}
}

func NewWordsFromFile(path string) *Words {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wl []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wl = append(wl, strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return NewWords(wl)
}

func (w *Words) ReducedList() []string {
	var list []string
	for _, val := range w.reducedList {
		if val != nil {
			list = append(list, *val)
		}
	}
	return list
}

func (w *Words) FullList() []string {
	var list []string
	for _, val := range w.fullList {
		list = append(list, val)
	}
	return list
}

func (w *Words) inList(gs string, r bool) bool {
	ugs := strings.ToUpper(gs)

	var wl []string

	if r {
		wl = w.ReducedList()
	} else {
		wl = w.FullList()
	}

	return common.StrInList(ugs, wl)
}

func (w *Words) InFullWords(gs string) bool {
	return w.inList(gs, false)
}

func (w *Words) InReducedWords(gs string) bool {
	return w.inList(gs, false)
}

func (w *Words) RandomWord() string {
	rand.Seed(time.Now().UnixNano())
	ri := rand.Intn(len(w.fullList))
	return w.fullList[ri]
}

/*
inc cases:
	true -> remove if the index selection is not met
	false -> remove if the index selection is met
*/
func (w *Words) ReduceByPosition(letter string, ind int, inc bool) {
	for i := 0; i < len(w.reducedList); i++ {
		match := false
		if letter == *w.reducedList[i] {
			match = true
		}

		if match != inc {
			w.reducedList[i] = nil
		}
	}
}

func (w *Words) ReduceByCount(letter string, count int, inc bool) {
	for i := 0; i < len(w.reducedList); i++ {
		var match bool
		if strings.Count(*w.reducedList[i], letter) >= count {
			match = true
		}

		if match != inc {
			w.reducedList[i] = nil
		}
	}
}
