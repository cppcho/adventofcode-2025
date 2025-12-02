package utils

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of all integers in a slice
func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// GCD returns the greatest common divisor of two integers
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two integers
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Contains checks if a slice contains a specific value
func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Map applies a function to each element of a slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := []T{}
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
