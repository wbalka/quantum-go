package maths

import (
	"math"
)

func logFactorial(n int) float64 {
	if n <= 1 {
		return 0.0
	}
	lnGamma, _ := math.Lgamma(float64(n + 1))
	return lnGamma
}

/*
Normalization constant for the hydrogenic radial wave function R_nl(r).
Using natural logarithms to avoid overflow on factorials
*/
func GetNormalizationConstant(n, l, Z int) float64 {
	ro := 2.0 * float64(Z) / float64(n)
	lnPart1 := 3.0 * math.Log(ro)
	lnNumerator := logFactorial(n - l - 1)
	lnDenominator := math.Log(2.0*float64(n)) + 3.0*logFactorial(n+l)

	lnNorm := lnPart1 + lnNumerator - lnDenominator
	return math.Exp(0.5 * lnNorm)
}
