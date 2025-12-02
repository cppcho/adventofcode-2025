package day02

import (
	"adventofcode2025/internal/utils"
	"fmt"
)

// Solve is the main entry point for day 2
func Solve(inputFile string) (interface{}, interface{}) {
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return nil, nil
	}

	part1Result := Part1(lines)
	part2Result := Part2(lines)

	return part1Result, part2Result
}

// Part1 solves part 1 of the puzzle
func Part1(lines []string) int {
	// TODO: Implement solution for part 1
	return 0
}

// Part2 solves part 2 of the puzzle
func Part2(lines []string) int {
	// TODO: Implement solution for part 2
	return 0
}
