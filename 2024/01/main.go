package main

import (
	"AdventOfCode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"codeberg.org/derat/advent-of-code/lib"
)

func main() {
	lines := lib.InputLines("2024/1")
	solvePuzzle(lines)
}

func solvePuzzle(lines []string) {
	left, right := parseLines(lines)

	// Part 1
	fmt.Println("Part 1: ", part1(left, right))

	// Part 2
	fmt.Println("Part 2: ", part2(left, right))
}

func parseLines(lines []string) ([]int, []int) {
	var col1 []int
	var col2 []int

	// Split each line into 2 numbers by consecutive whitespace
	for _, line := range lines {
		nums := strings.Fields(line)

		if len(nums) != 2 {
			fmt.Println("Invalid input", line)
			panic("Invalid input")
		}

		// Convert each number to an integer
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		// Add the numbers to the column arrays
		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	// Sort each column array
	sort.Ints(col1)
	sort.Ints(col2)

	return col1, col2
}

func part1(left []int, right []int) int {
	// Loop through the column arrays
	var sum int
	for i := 0; i < len(left); i++ {
		// Get the difference between the numbers at the index of each column array
		diff := right[i] - left[i]
		// Sum the differences using the ABS function
		sum += utils.Abs(diff)
	}

	// Return the sum
	return sum
}

func part2(left []int, right []int) int {
	var sum int
	var counts = getCounts(right)

	for i := 0; i < len(left); i++ {
		count := counts[left[i]]

		if count == 0 {
			continue
		}

		sum += count * left[i]
	}

	return sum
}

func getCounts(right []int) map[int]int {
	counts := make(map[int]int)

	for i := 0; i < len(right); i++ {
		counts[right[i]]++
	}

	return counts
}
