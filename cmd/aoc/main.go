package main

import (
	"fmt"
	"os"
	"strconv"

	"adventofcode2025/internal/day01"
	"adventofcode2025/internal/day02"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc <day>")
		fmt.Println("Example: aoc 1")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid day number: %s\n", os.Args[1])
		os.Exit(1)
	}

	switch day {
	case 1:
		runDay(day, day01.Solve)
	case 2:
		runDay(day, day02.Solve)
	// Add more days here as you implement them
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
		os.Exit(1)
	}
}

func runDay(day int, solve func(string) (any, any)) {
	inputFile := fmt.Sprintf("inputs/day%02d.txt", day)

	part1, part2 := solve(inputFile)

	fmt.Printf("Day %d Results:\n", day)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
