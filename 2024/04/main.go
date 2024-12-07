package main

import (
	"fmt"

	"codeberg.org/derat/advent-of-code/lib"
)

func main() {
	grid := lib.InputByteGrid("2024/4")
	solvePuzzle(grid)
}

func solvePuzzle(grid lib.ByteGrid) {
	wordsearch := WordSearch{
		grid: grid,
		w:    len(grid[0]),
		h:    len(grid),
	}

	fmt.Println(wordsearch)

	// Part 1
	fmt.Println("Part 1: ", part1(wordsearch))

	// Part 2
	fmt.Println("Part 2: ", part2())
}

func part1(search WordSearch) int {
	lookingFor := "XMAS"

}

func part2() int {
	return 0
}
