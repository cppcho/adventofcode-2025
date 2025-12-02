package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns all lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFile reads entire file content as a single string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadInts reads a file and returns all integers found (one per line)
func ReadInts(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var nums []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// SplitSections splits input by double newlines (common in AoC)
func SplitSections(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n\n")
}

// ParseIntsFromLine extracts all integers from a string (handles negative numbers)
func ParseIntsFromLine(line string) []int {
	var nums []int
	fields := strings.Fields(line)
	for _, field := range fields {
		// Remove common separators
		field = strings.Trim(field, ",:")
		if num, err := strconv.Atoi(field); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}
