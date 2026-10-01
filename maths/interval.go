package maths

type IntegralInterval struct {
	Start            float64
	End              float64
	Points           []float64
	Weights          []float64
	Values           []float64
	FirstDerivative  []float64
	SecondDerivative []float64
	Coefficients     []float64
	Hamiltonian      []float64
	IntegralValue    float64
}

type hamiltonian func(IntegralInterval, int) float64

func (interval *IntegralInterval) CalculateHamiltonian(fn hamiltonian) {
	interval.InterpolatingLegendrePolynomialCoefficients()
	interval.Hamiltonian = make([]float64, len(interval.Points))
	for i := range interval.Points {
		interval.Hamiltonian[i] = fn(*interval, i)
		//-0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
	}
}

func (interval *IntegralInterval) Integrate() {
	for i := range interval.Weights {
		r := interval.Points[i]
		integratedValue := r * r * interval.Values[i] * interval.Hamiltonian[i]
		interval.IntegralValue += integratedValue * interval.Weights[i]
	}
}
