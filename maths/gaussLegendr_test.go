package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestGaussLegendrePointsWeights(t *testing.T) {
	n := 3
	points, weights, err := GaussLegendrePointsWeights(n)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}
	fmt.Printf("Points: %v\n", points)
	fmt.Printf("Weights: %v\n", weights)
	fmt.Printf("Points length: %d\n", len(points))
	fmt.Printf("Weights length: %d\n", len(weights))
}

func TestLegendrePolynomialValue(t *testing.T) {
	n := 5
	x := 0.5
	expected := float64(float64(1)/float64(8)) * (63*math.Pow(x, 5) - 70*math.Pow(x, 3) + 15*x) // P_5(0.5)
	result := LegendrePolynomialValue(n, x)
	fmt.Printf("Expected polynomial value: %v, received:%v", expected, result)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestLegendrePolynomialValue2(t *testing.T) {
	n := 1
	x := 0.5
	expected := LegendrePolynomialValue(n, x)
	results, _, _ := LegendrePolynomialValuesWithDerivatives(n, x)
	result := results[n]
	fmt.Printf("Expected polynomial value: %v, received:%v", expected, result)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestLegendrePolynomialCoefficients(t *testing.T) {
	n := 5
	expected := []float64{0, 15.0 / 8.0, 0, -70.0 / 8.0, 0, 63.0 / 8.0} // Coefficients for P_5(x)
	result := LegendrePolynomialCoefficients(n)
	fmt.Printf("Expected coefficients: %v, received:%v", expected, result)
	for i := range result {
		if i > len(expected)-1 {
			if result[i] != 0 {
				t.Errorf("Expected coefficient at index %d to be 0, but got %v", i, result[i])
			}
			continue
		}
		if result[i] != expected[i] {
			t.Errorf("Expected coefficient at index %d to be %v, but got %v", i, expected[i], result[i])
		}
	}
}

func TestLegendreDerivativeCoefficients(t *testing.T) {
	n := 3
	expected := []float64{-1.5, 0, 7.5}
	result := DerivativeCoefficients(LegendrePolynomialCoefficients(n))
	fmt.Println(result)
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected coefficient at index %d to be %v, but got %v", i, expected[i], result[i])
		}
	}
}

func TestSecondLegendreDerivativeCoefficients(t *testing.T) {
	n := 4
	expected := []float64{-7.5, 0, 52.5}
	result := DerivativeCoefficients(DerivativeCoefficients(LegendrePolynomialCoefficients(n)))
	fmt.Println(result)
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected coefficient at index %d to be %v, but got %v", i, expected[i], result[i])
		}
	}
}

func TestLegendrePolynomialCoefficients2(t *testing.T) {
	n := 2
	expected := []float64{-0.5, 0, 1.5} // Coefficients for P_2(x)
	result := LegendrePolynomialCoefficients(n)
	fmt.Printf("Expected coefficients: %v, received:%v", expected, result)
	for i := range result {
		if i > len(expected)-1 {
			if result[i] != 0 {
				t.Errorf("Expected coefficient at index %d to be 0, but got %v", i, result[i])
			}
			continue
		}
		if result[i] != expected[i] {
			t.Errorf("Expected coefficient at index %d to be %v, but got %v", i, expected[i], result[i])
		}
	}
}

func TestInterpolatingLegendrePolynomialCoefficients(t *testing.T) {
	xMin := 2.0
	xMax := 10.0
	xpoints, weights, _ := GaussLegendrePointsWeights(40)
	values := make([]float64, len(xpoints))
	for j := range xpoints {
		xpoints[j] = (xMax-xMin)/2.0*xpoints[j] + (xMax+xMin)/2.0
	}
	alpha := 1.0 / math.Sqrt(math.Pi)
	for i := range values {
		values[i] = alpha * math.Exp(xpoints[i])
	}
	interval := IntegralInterval{
		Start:   xMin,
		End:     xMax,
		Points:  xpoints,
		Weights: weights,
		Values:  values,
	}
	interval.InterpolatingLegendrePolynomialCoefficients()
	returnedValue, _, _ := CalculateValueUsingCoefficients(interval.Coefficients, xpoints[6], xMin, xMax)
	fmt.Printf("Received value: %f, expected %f", returnedValue, values[6])
	delta := 0.00000000001
	if math.Abs(returnedValue-values[6]) > delta {
		t.Errorf("Expected value: %.17g, but got: %.17g", values[6], returnedValue)
	}
}
func TestTryToGetFirstDerivative(t *testing.T) {
	xMin := 2.0
	xMax := 4.0
	gLPoints := 33
	xpoints, weights, _ := GaussLegendrePointsWeights(gLPoints)
	values := make([]float64, len(xpoints))
	for j := range xpoints {
		xpoints[j] = (xMax-xMin)/2.0*xpoints[j] + (xMax+xMin)/2.0
	}
	alpha := 1.0 / math.Sqrt(math.Pi)
	for i := range values {
		values[i] = alpha * math.Exp(-xpoints[i])
	}
	interval := IntegralInterval{
		Start:            xMin,
		End:              xMax,
		Points:           xpoints,
		Weights:          weights,
		Values:           values,
		FirstDerivative:  make([]float64, gLPoints),
		SecondDerivative: make([]float64, gLPoints),
		Coefficients:     make([]float64, gLPoints),
	}
	interval.InterpolatingLegendrePolynomialCoefficients()
	//value, firstDerivative, secondDerivative := CalculateValueUsingCoefficients(coeffs, x, xMin, xMax)
	//derivativeValue := TryToGetFirstDerivative(coeffs, x, xMin, xMax)
	fmt.Printf("The real value     is: %.17g \n", interval.Values[32])
	//fmt.Printf("The interpolated v was %.17g \n", value)
	fmt.Printf("The 1st derivative was %.17g \n", interval.FirstDerivative[32])
	//fmt.Printf("The old 1st deriv  was %.17g \n", derivativeValue)
	fmt.Printf("The 2nd derivative was %.17g \n", interval.SecondDerivative[32])
}

func TestLPPvalue(t *testing.T) {
	oldValue := LegendrePolynomialValue(3, -0.5)
	value, firstDerivative, secondDerivative := LegendrePolynomialValuesWithDerivatives(3, -0.5)
	fmt.Printf("The old value was %.17g \n", oldValue)
	fmt.Printf("The value was %.17g \n", value)
	fmt.Printf("The 1st derivative was  %.17g \n", firstDerivative)
	fmt.Printf("The 2nd derivative was  %.17g \n", secondDerivative)

}
