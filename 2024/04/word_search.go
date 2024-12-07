package main

import "codeberg.org/derat/advent-of-code/lib"

type LetterPos struct {
	letter string
	x      int
	y      int
}

type Word struct {
	value string
	path  []LetterPos
}

type WordSearch struct {
	grid lib.ByteGrid
	w    int
	h    int
}

func (ws WordSearch) checkPos(x int, y int, word Word) []Word {

}

func (ws WordSearch) FindWord(word string) []Word {
	var foundWords []Word

	for y := 0; y < ws.h; y++ {
		for x := 0; x < ws.w; x++ {
			found := ws.checkPos(x, y, Word{value: word})

			if len(found) > 0 {
				foundWords = append(foundWords, found...)
			}
		}
	}

	return foundWords
}
