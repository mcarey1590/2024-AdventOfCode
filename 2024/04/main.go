package main

import (
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
	letter string
	x      int
	y      int
}

type Word struct {
	value string
	shape [][]int // 1 = letter, 0 = don't care
	path  []LetterPos
}

type WordSearch struct {
	directions []Direction
	grid       lib.ByteGrid
	w          int
	h          int
}

func (ws WordSearch) checkDirection(dir Direction, word Word) *Word {
	lastPath := word.path[len(word.path)-1]
	x := lastPath.x + dir.x
	y := lastPath.y + dir.y

	if x < 0 || x >= ws.w || y < 0 || y >= ws.h {
		return nil
	}

	lookingFor := word.value[len(word.path)]
	if lookingFor != ws.grid[y][x] {
		return nil
	}

	newWord := Word{value: word.value, path: append(word.path, LetterPos{letter: string(lookingFor), x: x, y: y})}

	if len(newWord.path) == len(word.value) {
		return &newWord
	}

	return ws.checkDirection(dir, newWord)

}

func (ws WordSearch) findWordStartingAt(x int, y int, word Word) []Word {
	if ws.grid[y][x] != word.value[0] {
		return []Word{}
	}

	var foundWords []Word
	for _, dir := range ws.directions {
		found := ws.checkDirection(dir, Word{value: word.value, path: []LetterPos{{letter: string(word.value[0]), x: x, y: y}}})
		if found != nil {
			fmt.Println("Found words: ", found)
			foundWords = append(foundWords, *found)
		}
	}

	return foundWords
}

func (ws WordSearch) findWordShapeStartingAt(x int, y int, word Word) []Word {

}

func main() {
	grid := lib.NewByteGridString("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
	//grid := lib.InputByteGrid("2024/4")
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
		value: "XMAS",
		shape: [][]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
		},
	}

	for y := 0; y < ws.h; y++ {
		for x := 0; x < ws.w; x++ {
			fmt.Println("Checking: ", x, y)
			found := ws.findWordShapeStartingAt(x, y, word)

			if len(found) > 0 {
				fmt.Println("Found words - base: ", found)
				foundWords = append(foundWords, found...)
			}
		}
	}

	return len(foundWords)
}
