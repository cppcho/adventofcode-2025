# Advent of Code 2025

Solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Project Structure

```
adventofcode-2025/
├── cmd/aoc/          # CLI runner
├── internal/         # Solution packages
│   ├── day01/       # Day 1 solution
│   ├── day02/       # Day 2 solution
│   └── utils/       # Shared utilities
└── inputs/          # Puzzle inputs (not committed to git)
```

## Usage

### Running Solutions

Run a specific day's solution:

```bash
go run cmd/aoc/main.go <day>
```

Example:
```bash
go run cmd/aoc/main.go 1
```

### Running Tests

Run tests for all days:
```bash
go test ./...
```

Run tests for a specific day:
```bash
go test ./internal/day01
```

### Adding a New Day

1. Copy the template from `internal/day01/`
2. Create a new directory `internal/dayXX/`
3. Implement the `Solve()`, `Part1()`, and `Part2()` functions
4. Add your puzzle input to `inputs/dayXX.txt`
5. Update the CLI runner in `cmd/aoc/main.go` if needed

## Development

Build the CLI:
```bash
go build -o aoc cmd/aoc/main.go
./aoc 1
```

Format code:
```bash
go fmt ./...
```

Run linter (requires golangci-lint):
```bash
golangci-lint run
```

## Notes

- Puzzle inputs are not committed to git per Advent of Code's request
- Each day is its own package for clean separation and testability
- The `utils` package contains shared helper functions for input parsing and common algorithms
