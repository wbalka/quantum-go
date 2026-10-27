package maths

type IntegralArea struct {
	StartRadial             float64
	EndRadial               float64
	StartAngular            float64
	EndAngular              float64
	RadialIntervalSize      float64
	AngularIntervalSize     float64
	RadialPoints            []float64
	RadialWeights           []float64
	AngularPoints           []float64
	AngularWeights          []float64
	Values                  [][]float64
	RadialFirstDerivative   [][]float64
	RadialSecondDerivative  [][]float64
	AngularFirstDerivative  [][]float64
	AngularSecondDerivative [][]float64
	KineticEnergy           [][]float64
	Potential               [][]float64
	IntegralKinetic         float64
	IntegralPotential       float64
}

type RadialInterval struct {
	Start            float64
	End              float64
	Points           []float64
	Weights          []float64
	Values           []float64
	FirstDerivative  []float64
	SecondDerivative []float64
	Hamiltonian      []float64
	IntegralValue    float64
}

type InterpolationInterval struct {
	Start            float64
	End              float64
	Points           []float64
	Weights          []float64
	Values           []float64
	FirstDerivative  []float64
	SecondDerivative []float64
	Coefficients     []float64
}

type radialHamiltonian func(RadialInterval, int) float64

type sphericalHamiltonian2d func(IntegralArea, int, int) (float64, float64)

func (s *RadialInterval) calculateDerivatives() {
	interpolationInterval := InterpolationInterval{
		Start:   s.Start,
		End:     s.End,
		Points:  s.Points,
		Weights: s.Weights,
		Values:  s.Values,
	}
	interpolationInterval.CalculateInterpolatingPolynomialCoefficientsAndFunctionDerivatives()
	s.FirstDerivative = interpolationInterval.FirstDerivative
	s.SecondDerivative = interpolationInterval.SecondDerivative
}

func (interval *RadialInterval) CalculateHamiltonian(fn radialHamiltonian) {
	interval.calculateDerivatives()
	interval.Hamiltonian = make([]float64, len(interval.Points))
	for i := range interval.Points {
		interval.Hamiltonian[i] = fn(*interval, i)
		//-0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
	}
}

func (interval *RadialInterval) Integrate() {
	for i := range interval.Weights {
		r := interval.Points[i]
		integratedValue := r * r * interval.Values[i] * interval.Hamiltonian[i]
		interval.IntegralValue += integratedValue * interval.Weights[i]
	}
}

func (s *IntegralArea) calculateDerivatives() {
	s.RadialFirstDerivative = make([][]float64, len(s.RadialPoints))
	s.RadialSecondDerivative = make([][]float64, len(s.RadialPoints))
	s.AngularFirstDerivative = make([][]float64, len(s.RadialPoints))
	s.AngularSecondDerivative = make([][]float64, len(s.RadialPoints))
	for r := range s.RadialPoints {
		s.RadialFirstDerivative[r] = make([]float64, len(s.AngularPoints))
		s.RadialSecondDerivative[r] = make([]float64, len(s.AngularPoints))
	}
	for u := range s.AngularPoints {
		values := make([]float64, len(s.RadialPoints))
		for r := range s.Values {
			values[r] = s.Values[r][u]
		}
		interpolationInterval := InterpolationInterval{
			Start:   s.StartRadial,
			End:     s.EndRadial,
			Points:  s.RadialPoints,
			Weights: s.RadialWeights,
			Values:  values,
		}
		interpolationInterval.CalculateInterpolatingPolynomialCoefficientsAndFunctionDerivatives()
		for r := range s.RadialPoints {
			s.RadialFirstDerivative[r][u] = interpolationInterval.FirstDerivative[r]
			s.RadialSecondDerivative[r][u] = interpolationInterval.SecondDerivative[r]
		}
	}
	for r := range s.RadialPoints {
		interpolationInterval := InterpolationInterval{
			Start:   s.StartAngular,
			End:     s.EndAngular,
			Points:  s.AngularPoints,
			Weights: s.AngularWeights,
			Values:  s.Values[r],
		}
		interpolationInterval.CalculateInterpolatingPolynomialCoefficientsAndFunctionDerivatives()
		s.AngularFirstDerivative[r] = interpolationInterval.FirstDerivative
		s.AngularSecondDerivative[r] = interpolationInterval.SecondDerivative
	}
}

func (interval *IntegralArea) CalculateHamiltonian(fn sphericalHamiltonian2d) {
	interval.calculateDerivatives()
	interval.KineticEnergy = make([][]float64, len(interval.RadialPoints))
	interval.Potential = make([][]float64, len(interval.RadialPoints))
	for r := range interval.RadialPoints {
		interval.KineticEnergy[r] = make([]float64, len(interval.AngularPoints))
		interval.Potential[r] = make([]float64, len(interval.AngularPoints))
		for u := range interval.AngularPoints {
			interval.KineticEnergy[r][u], interval.Potential[r][u] = fn(*interval, r, u)
		}
	}
}

func (interval *IntegralArea) Integrate() {
	for i := range interval.RadialWeights {
		for j := range interval.AngularWeights {
			r := interval.RadialPoints[i]
			integratedKinetic := r * r * interval.Values[i][j] * interval.KineticEnergy[i][j]
			integratedPotential := r * r * interval.Values[i][j] * interval.Potential[i][j]
			interval.IntegralKinetic += integratedKinetic * interval.RadialWeights[i] * interval.AngularWeights[j]
			interval.IntegralPotential += integratedPotential * interval.RadialWeights[i] * interval.AngularWeights[j]
		}
	}
}
