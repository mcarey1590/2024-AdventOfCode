package main

import (
	"AdventOfCode/utils"
	"codeberg.org/derat/advent-of-code/lib"
	"fmt"
)

func main() {
	lines := lib.InputLines("2024/2")
	solvePuzzle(lines)
}

func solvePuzzle(lines []string) {
	reportsData := parseLines(lines)

	// Part 1
	fmt.Println("Part 1: ", part1(reportsData))

	// Part 2
	fmt.Println("Part 2: ", part2(reportsData))
}

func parseLines(lines []string) [][]int {
	var nums [][]int
	for _, line := range lines {
		nums = append(nums, lib.ExtractInts(line))
	}

	return nums
}

func part1(reportsData [][]int) int {
	return determineSafeReports(reportsData, 0)
}

func part2(reportsData [][]int) int {
	return determineSafeReports(reportsData, 1)
}

func determineSafeReports(reportsData [][]int, tolerance int) int {
	safeReports := 0

	for _, report := range reportsData {
		safeReport := determineUnsafeLevels(report, tolerance)

		if safeReport {
			safeReports++
		} else {
			//fmt.Println("Retry without first level")
			safeReport = determineUnsafeLevels(report[1:], tolerance-1)

			if safeReport {
				safeReports++
			}
		}
	}

	return safeReports
}

func determineUnsafeLevels(report []int, tolerance int) bool {
	unsafeLevels := 0
	prevLevel := report[0]
	dir := 0
	for _, nextLevel := range report[1:] {
		if unsafeLevels > tolerance {
			break
		}

		levelDiff := nextLevel - prevLevel

		if levelDiff == 0 || utils.Abs(levelDiff) > 3 {
			unsafeLevels++
			continue
		}

		currDir := 0
		if levelDiff > 0 {
			currDir = 1
		} else {
			currDir = -1
		}

		if dir != 0 && currDir != dir {
			unsafeLevels++
			continue
		}

		if dir == 0 {
			dir = currDir
		}

		prevLevel = nextLevel
	}
	safe := unsafeLevels <= tolerance

	//if !safe {
	//	fmt.Println(report)
	//}

	return safe
}
