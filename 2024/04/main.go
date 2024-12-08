package main

import (
	"AdventOfCode/utils"
	"fmt"

	"codeberg.org/derat/advent-of-code/lib"
)

type Direction struct {
	x int
	y int
}

var (
	left      = Direction{x: -1, y: 0}
	right     = Direction{x: 1, y: 0}
	up        = Direction{x: 0, y: -1}
	down      = Direction{x: 0, y: 1}
	upLeft    = Direction{x: -1, y: -1}
	upRight   = Direction{x: 1, y: -1}
	downLeft  = Direction{x: -1, y: 1}
	downRight = Direction{x: 1, y: 1}
)

type LetterPos struct {
	letter byte
	x      int
	y      int
}

type Word struct {
	value string
	path  []LetterPos
}

type WordSearch struct {
	directions []Direction
	grid       lib.ByteGrid
	w          int
	h          int
}

func main() {
	grid := lib.InputByteGrid("2024/4")
	solvePuzzle(grid)
}

func solvePuzzle(grid lib.ByteGrid) {
	wordsearch := WordSearch{
		directions: []Direction{left, right, up, down, upLeft, upRight, downLeft, downRight},
		grid:       grid,
		w:          len(grid[0]),
		h:          len(grid),
	}

	// Part 1
	fmt.Println("Part 1: ", part1(wordsearch))

	// Part 2
	fmt.Println("Part 2: ", part2(wordsearch))
}

func part1(ws WordSearch) int {
	var foundWords []Word
	word := Word{value: "XMAS"}

	for y := 0; y < ws.h; y++ {
		for x := 0; x < ws.w; x++ {
			found := ws.findWordStartingAt(x, y, word)

			if len(found) > 0 {
				foundWords = append(foundWords, found...)
			}
		}
	}

	return len(foundWords)
}

func part2(ws WordSearch) int {
	var foundWords []Word
	word := Word{
		value: "MAS",
	}

	for y := 0; y < ws.h; y++ {
		for x := 0; x < ws.w; x++ {
			found := ws.findIntersectingWords(x, y, word)

			if found != nil {
				foundWords = append(foundWords, *found)
			}
		}
	}

	return len(foundWords)
}

func (ws WordSearch) findWordStartingAt(x int, y int, word Word) []Word {
	if ws.grid[y][x] != word.value[0] {
		return []Word{}
	}

	var foundWords []Word
	for _, dir := range ws.directions {
		found := ws.checkDirection(dir, Word{value: word.value, path: []LetterPos{{letter: word.value[0], x: x, y: y}}})
		if found != nil {
			foundWords = append(foundWords, *found)
		}
	}

	return foundWords
}

func (ws WordSearch) getNextLetter(currentPos LetterPos, dir Direction) *LetterPos {
	x := currentPos.x + dir.x
	y := currentPos.y + dir.y

	if x < 0 || x >= ws.w || y < 0 || y >= ws.h {
		return nil
	}

	return &LetterPos{letter: ws.grid[y][x], x: x, y: y}
}

func (ws WordSearch) checkDirection(dir Direction, word Word) *Word {
	lastPath := word.path[len(word.path)-1]
	nextLetter := ws.getNextLetter(lastPath, dir)

	lookingFor := word.value[len(word.path)]
	if nextLetter == nil || lookingFor != nextLetter.letter {
		return nil
	}

	newWord := Word{value: word.value, path: append(word.path, *nextLetter)}

	if len(newWord.path) == len(word.value) {
		return &newWord
	}

	return ws.checkDirection(dir, newWord)

}

func (ws WordSearch) findIntersectingWords(x int, y int, word Word) *Word {
	wordLength := len(word.value)
	lib.Assertf(wordLength%2 == 1, "Word value must be an odd number of characters")
	distanceFromCenter := wordLength / 2

	centerLetter := LetterPos{letter: word.value[distanceFromCenter], x: x, y: y}

	if ws.grid[y][x] != centerLetter.letter {
		return nil
	}

	// Check diagonals
	leftToRight := make([]LetterPos, wordLength)
	leftToRight[distanceFromCenter] = centerLetter
	rightToLeft := make([]LetterPos, wordLength)
	rightToLeft[distanceFromCenter] = centerLetter

	for i := 1; i <= distanceFromCenter; i++ {
		left := distanceFromCenter - i
		right := distanceFromCenter + i

		if next := ws.getNextLetter(leftToRight[left+1], upLeft); next != nil {
			leftToRight[left] = *next
		}

		if next := ws.getNextLetter(leftToRight[right-1], downRight); next != nil {
			leftToRight[right] = *next
		}

		if next := ws.getNextLetter(rightToLeft[left+1], upRight); next != nil {
			rightToLeft[left] = *next
		}

		if next := ws.getNextLetter(rightToLeft[right-1], downLeft); next != nil {
			rightToLeft[right] = *next
		}
	}

	return checkIntersectionWords(&word, leftToRight, rightToLeft)
}

func getWorkFromLetters(letters []LetterPos) string {
	word := ""
	for _, letter := range letters {
		word += string(letter.letter)
	}
	return word
}

func checkIntersectionWords(word *Word, toRight []LetterPos, toLeft []LetterPos) *Word {
	left := getWorkFromLetters(toLeft)
	right := getWorkFromLetters(toRight)

	if left != word.value && utils.Reverse(left) != word.value {
		return nil
	}

	if right != word.value && utils.Reverse(right) != word.value {
		return nil
	}

	word.path = append(toLeft, toRight...)

	return word
}
