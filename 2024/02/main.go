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

	lib.Assert(len(reportsData) == 1000)

	// Part 1
	safeReports, unsafeReports := part1(reportsData)
	part1Answer := len(safeReports)
	fmt.Println("Part 1: ", part1Answer)

	// Part 2
	safeWithTolerance, unsafeWithTolerance := part2(unsafeReports)
	part2Answer := part1Answer + len(safeWithTolerance)
	fmt.Println("Part 2: ", part2Answer)

	lib.Assert(part2Answer+len(unsafeWithTolerance) == 1000)
}

func parseLines(lines []string) [][]int {
	var nums [][]int
	for _, line := range lines {
		lineNums := lib.ExtractInts(line)
		nums = append(nums, lineNums)
	}

	return nums
}

func part1(reportsData [][]int) ([][]int, [][]int) {
	return conductSafetyTest(reportsData, false)
}

func part2(unsafeReports [][]int) ([][]int, [][]int) {
	return conductSafetyTest(unsafeReports, true)
}

func conductSafetyTest(reportsData [][]int, checkWithTolerance bool) ([][]int, [][]int) {
	safeReports := make([][]int, 0)
	unsafeReports := make([][]int, 0)

	for _, report := range reportsData {
		safeReport := isReportSafe(report)

		if !safeReport && checkWithTolerance {
			for i := 0; i < len(report); i++ {
				if isReportSafe(utils.RemoveIndex(report, i)) {
					safeReport = true
					break
				}
			}
		}

		if safeReport {
			safeReports = append(safeReports, report)
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	return safeReports, unsafeReports
}

func isReportSafe(report []int) bool {
	prevLevel := report[0]
	dir := 0

	for _, nextLevel := range report[1:] {
		nextLevelIsSafe, currDir := isNextLevelSafe(prevLevel, nextLevel, dir)

		if !nextLevelIsSafe {
			return false
		}

		if dir == 0 {
			dir = currDir
		}

		prevLevel = nextLevel
	}

	return true
}

func isNextLevelSafe(prevLevel int, nextLevel int, dir int) (bool, int) {
	levelDiff := nextLevel - prevLevel

	currDir := 0
	if levelDiff > 0 {
		currDir = 1
	} else {
		currDir = -1
	}

	if dir != 0 && currDir != dir {
		return false, dir
	}

	if levelDiff == 0 || utils.Abs(levelDiff) > 3 {
		return false, dir
	}

	return true, currDir
}
