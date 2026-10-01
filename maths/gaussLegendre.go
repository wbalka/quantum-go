package maths

import (
	"fmt"
	"math"
)

func GaussLegendrePointsWeights(n int) ([]float64, []float64, error) {
	if n <= 0 {
		return nil, nil, fmt.Errorf("n must be a positive integer")
	}

	points := make([]float64, n)
	weights := make([]float64, n)

	m := (n + 1) / 2
	eps := 1e-15

	for i := range m {
		z := math.Cos(math.Pi * (float64(i) + 0.75) / (float64(n) + 0.5))
		var p1, pp float64
		for {
			p1 = 1.0 // P_0(z)
			p2 := 0.0
			for k := 1; k <= n; k++ {
				p3 := p2
				p2 = p1
				p1 = ((float64(2*k-1) * z * p2) - (float64(k-1) * p3)) / float64(k)
			}
			pp = float64(n) * (z*p1 - p2) / (z*z - 1.0)
			z1 := z
			z = z1 - p1/pp
			if math.Abs(z-z1) < eps {
				break
			}
		}
		points[i] = -z
		points[n-1-i] = z
		weight := 2.0 / ((1.0 - z*z) * pp * pp)
		weights[i] = weight
		weights[n-1-i] = weight
	}

	return points, weights, nil
}

func LegendrePolynomialValue(n int, x float64) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	return (float64(2*(n-1)+1)*x*LegendrePolynomialValue(n-1, x)/float64(n) - float64(n-1)*LegendrePolynomialValue(n-2, x)/float64(n))
}

func LegendrePolynomialValuesWithDerivatives(n int, x float64) (lp, dlp, ddlp []float64) {
	lp, dlp, ddlp = make([]float64, n+1), make([]float64, n+1), make([]float64, n+1)
	lp[0] = 1
	if n > 0 {
		lp[1], dlp[1] = x, 1
	}
	for i := 1; i < n; i++ {
		lp[i+1] = (float64(2*i+1)*x*lp[i] - float64(i)*lp[i-1]) / float64(i+1)
		dlp[i+1] = dlp[i-1] + float64(2*i+1)*lp[i]
		ddlp[i+1] = ddlp[i-1] + float64(2*i+1)*dlp[i]
	}
	return
}

func LegendrePolynomialCoefficients(n int) []float64 {
	coefficients := make([]float64, 100)
	if n == 0 {
		coefficients[0] = 1.0
		return coefficients
	}
	if n == 1 {
		coefficients[0] = 0.0
		coefficients[1] = 1.0
		return coefficients
	}
	recursiveCoefficients := LegendrePolynomialCoefficients(n - 1)
	recursiveCoefficients2 := LegendrePolynomialCoefficients(n - 2)
	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			coefficients[i] = -float64(n-1) * recursiveCoefficients2[i] / float64(n)
		case n:
			coefficients[i] = float64(2*(n-1)+1) * recursiveCoefficients[i-1] / float64(n)
		default:
			coefficients[i] = (float64(2*(n-1)+1)*recursiveCoefficients[i-1] - float64(n-1)*recursiveCoefficients2[i]) / float64(n)
		}
	}
	return coefficients
}

func DerivativeCoefficients(coefficients []float64) []float64 {
	n := 0
	for i := range coefficients {
		if coefficients[i] != 0 {
			n = i + 1
		}
	}
	dvCoeffs := make([]float64, n-1)
	for i := range dvCoeffs {
		dvCoeffs[i] = coefficients[i+1] * float64(i+1)
	}
	return dvCoeffs
}

func (interval *IntegralInterval) InterpolatingLegendrePolynomialCoefficients() error {
	xpoints := interval.Points
	values := interval.Values
	xMin := interval.Start
	xMax := interval.End
	interval.Coefficients = make([]float64, len(xpoints))
	interval.FirstDerivative = make([]float64, len(xpoints))
	interval.SecondDerivative = make([]float64, len(xpoints))
	_, weights, error := GaussLegendrePointsWeights(len(xpoints))
	if error != nil {
		return error
	}
	normalizedX := make([]float64, len(xpoints))
	for i, x := range xpoints {
		normalizedX[i] = (2.0*x - (xMin + xMax)) / (xMax - xMin)
	}
	valuesMatrix := make([][]float64, len(normalizedX))
	dMatrix := make([][]float64, len(normalizedX))
	ddMatrix := make([][]float64, len(normalizedX))
	for j := range normalizedX {
		valuesMatrix[j], dMatrix[j], ddMatrix[j] = LegendrePolynomialValuesWithDerivatives(len(interval.Coefficients)-1, normalizedX[j])
	}
	for i := range interval.Coefficients {
		c := 0.0
		for j := range normalizedX {
			c += weights[j] * values[j] * valuesMatrix[j][i]
		}
		interval.Coefficients[i] = c * float64(2.0*i+1.0) / 2.0
	}
	scale := 2 / (xMax - xMin)

	for i := range normalizedX {
		firstDv := 0.0
		secondDv := 0.0
		for j, co := range interval.Coefficients {
			firstDv += co * dMatrix[i][j]
			secondDv += co * ddMatrix[i][j]
		}
		firstDv *= scale
		secondDv *= scale * scale
		interval.FirstDerivative[i] = firstDv
		interval.SecondDerivative[i] = secondDv
	}
	return nil
}

func CalculateValueUsingCoefficients(coeffs []float64, point float64, xMin float64, xMax float64) (value, firstDv, secondDv float64) {
	scale := 2 / (xMax - xMin)
	normalizedX := (2.0*point - (xMin + xMax)) / (xMax - xMin)
	lp, dlp, ddlp := LegendrePolynomialValuesWithDerivatives(len(coeffs)-1, normalizedX)
	for i, co := range coeffs {
		value += co * lp[i]
		firstDv += co * dlp[i]
		secondDv += co * ddlp[i]
	}
	firstDv *= scale
	secondDv *= scale * scale
	return
}

func TryToGetFirstDerivative(coeffs []float64, point float64, xMin float64, xMax float64) float64 {
	normalizedX := float64((2.0*point - (xMin + xMax)) / (xMax - xMin))
	value := 0.0
	for i := range len(coeffs) - 1 {
		coeff := coeffs[i+1] * (2 / (xMax - xMin))
		dvCoeffs := DerivativeCoefficients(LegendrePolynomialCoefficients(i + 1))
		for j, c := range dvCoeffs {
			value += c * math.Pow(normalizedX, float64(j)) * coeff
		}
	}
	return value
}
