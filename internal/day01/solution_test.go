package day01

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "example 1",
			input: []string{
				// TODO: Add example input
			},
			expected: 0, // TODO: Add expected result
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part1(tt.input)
			if result != tt.expected {
				t.Errorf("Part1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "example 1",
			input: []string{
				// TODO: Add example input
			},
			expected: 0, // TODO: Add expected result
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part2(tt.input)
			if result != tt.expected {
				t.Errorf("Part2() = %v, want %v", result, tt.expected)
			}
		})
	}
}
