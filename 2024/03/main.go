package main

import (
	"AdventOfCode/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"codeberg.org/derat/advent-of-code/lib"
)

func main() {
	lines := lib.InputLines("2024/3")
	solvePuzzle(lines)
}

type InstructionType string

const (
	Multiply InstructionType = "mul"
	Do       InstructionType = "do"
	Dont     InstructionType = "don't"
)

type Instruction struct {
	x               int
	y               int
	instructionType InstructionType
}

func (i Instruction) multiply() int {
	return i.x * i.y
}

func solvePuzzle(lines []string) {
	instructions := parseLines(lines)

	// Part 1
	fmt.Println("Part 1: ", part1(instructions))

	// Part 2
	fmt.Println("Part 2: ", part2(instructions))
}

func parseLines(lines []string) []Instruction {
	pattern := regexp.MustCompile(`(?m)((mul)\(((\d{1,3}),(\d{1,3}))\))|((do)\(\))|((don't)\(\))`)

	instructions := make([]Instruction, 0)

	for _, line := range lines {
		matches := pattern.FindAllStringSubmatch(line, -1)

		lib.Assertf(len(matches) > 0, "line contains no matches: %v", line)

		for _, match := range matches {
			sanitized := utils.RemoveEmptyStrings(match)
			lib.Assertf(len(sanitized) >= 3, "Invalid match: %v", sanitized)

			instructionType := InstructionType(sanitized[2])
			var instruction Instruction

			switch instructionType {
			case Multiply:
				x, _ := strconv.Atoi(sanitized[4])
				y, _ := strconv.Atoi(sanitized[5])
				instruction = Instruction{
					instructionType: instructionType,
					x:               x,
					y:               y,
				}
			case Do:
				instruction = Instruction{
					instructionType: Do,
				}
			case Dont:
				instruction = Instruction{
					instructionType: Dont,
				}
			default:
				log.Panicf("%v is not a valid instruction type. \nMatch: %v", instructionType, strings.Join(match, "|"))
			}

			instructions = append(instructions, instruction)
		}
	}

	return instructions
}

func part1(instructions []Instruction) int {
	result := 0
	for _, instruction := range instructions {
		result += instruction.multiply()
	}
	return result
}

func part2(instructions []Instruction) int {
	disabled := false
	result := 0

	for _, instruction := range instructions {
		switch instruction.instructionType {
		case Multiply:
			if !disabled {
				result += instruction.multiply()
			}
		case Do:
			disabled = false
		case Dont:
			disabled = true
		}
	}

	return result
}
