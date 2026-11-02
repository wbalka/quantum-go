package maths

import (
	"math"
)

func factorial(n int) int {
	if n < 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func NormalizationConstant(n, l, Z int) float64 {
	ro := (2.0 * float64(Z) / float64(n))
	part1 := ro * ro * ro
	part2 := float64(factorial(n-l-1)) / (2 * float64(n) * math.Pow(float64(factorial(n+l)), 3))
	return math.Sqrt(part1 * part2)
}
