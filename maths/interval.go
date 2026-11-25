package maths

import (
	"fmt"
	"math"

	"golang.org/x/sync/errgroup"
)

type Interval interface {
	Integrate()
}

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

type IntegralVolume struct {
	StartRadial               float64
	EndRadial                 float64
	StartAngular              float64
	EndAngular                float64
	RadialIntervalSize        float64
	AngularIntervalSize       float64
	StartAzimuthal            float64
	EndAzimuthal              float64
	AzimuthalIntervalSize     float64
	RadialPoints              []float64
	RadialWeights             []float64
	AngularPoints             []float64
	AngularWeights            []float64
	AzimuthalPoints           []float64
	AzimuthalWeights          []float64
	Values                    [][][]float64
	RadialFirstDerivative     [][][]float64
	RadialSecondDerivative    [][][]float64
	AngularFirstDerivative    [][][]float64
	AngularSecondDerivative   [][][]float64
	AzimuthalFirstDerivative  [][][]float64
	AzimuthalSecondDerivative [][][]float64
	KineticEnergy             [][][]float64
	Potential                 [][][]float64
	IntegralKinetic           float64
	IntegralPotential         float64
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

type RadialHamiltonian func(RadialInterval, int) float64

type SphericalHamiltonian2d func(IntegralArea, int, int) (float64, float64)

type SphericalHamiltonian3d func(IntegralVolume, int, int, int) (float64, float64)

type Calculation func(Interval) error

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

func (interval *RadialInterval) CalculateHamiltonian(fn RadialHamiltonian) {
	interval.calculateDerivatives()
	interval.Hamiltonian = make([]float64, len(interval.Points))
	for i := range interval.Points {
		interval.Hamiltonian[i] = fn(*interval, i)
		//-0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
	}
}

func (interval *RadialInterval) Integrate() {
	interval.IntegralValue = 0
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

func (interval *IntegralArea) CalculateHamiltonian(fn SphericalHamiltonian2d) {
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
	interval.IntegralKinetic = 0
	interval.IntegralPotential = 0
	for i := range interval.RadialWeights {
		r := interval.RadialPoints[i]
		for j := range interval.AngularWeights {
			th := interval.AngularPoints[j]
			integratedKinetic := r * r * math.Sin(th) * interval.Values[i][j] * interval.KineticEnergy[i][j]
			integratedPotential := r * r * math.Sin(th) * interval.Values[i][j] * interval.Potential[i][j]
			interval.IntegralKinetic += integratedKinetic * interval.RadialWeights[i] * interval.AngularWeights[j]
			interval.IntegralPotential += integratedPotential * interval.RadialWeights[i] * interval.AngularWeights[j]
		}
	}
}

func (s *IntegralVolume) calculateDerivatives() {
	s.RadialFirstDerivative = make([][][]float64, len(s.RadialPoints))
	s.RadialSecondDerivative = make([][][]float64, len(s.RadialPoints))
	s.AngularFirstDerivative = make([][][]float64, len(s.RadialPoints))
	s.AngularSecondDerivative = make([][][]float64, len(s.RadialPoints))
	s.AzimuthalFirstDerivative = make([][][]float64, len(s.RadialPoints))
	s.AzimuthalSecondDerivative = make([][][]float64, len(s.RadialPoints))
	for r := range s.RadialPoints {
		s.RadialFirstDerivative[r] = make([][]float64, len(s.AngularPoints))
		s.RadialSecondDerivative[r] = make([][]float64, len(s.AngularPoints))
		s.AngularFirstDerivative[r] = make([][]float64, len(s.AngularPoints))
		s.AngularSecondDerivative[r] = make([][]float64, len(s.AngularPoints))
		s.AzimuthalFirstDerivative[r] = make([][]float64, len(s.AngularPoints))
		s.AzimuthalSecondDerivative[r] = make([][]float64, len(s.AngularPoints))
		for u := range s.AngularPoints {
			s.RadialFirstDerivative[r][u] = make([]float64, len(s.AzimuthalPoints))
			s.RadialSecondDerivative[r][u] = make([]float64, len(s.AzimuthalPoints))
			s.AngularFirstDerivative[r][u] = make([]float64, len(s.AzimuthalPoints))
			s.AngularSecondDerivative[r][u] = make([]float64, len(s.AzimuthalPoints))
		}
	}
	for r := range s.RadialPoints {
		for u := range s.AngularPoints {
			interpolationInterval := InterpolationInterval{
				Start:   s.StartAzimuthal,
				End:     s.EndAzimuthal,
				Points:  s.AzimuthalPoints,
				Weights: s.AzimuthalWeights,
				Values:  s.Values[r][u],
			}
			interpolationInterval.CalculateInterpolatingPolynomialCoefficientsAndFunctionDerivatives()
			s.AzimuthalFirstDerivative[r][u] = interpolationInterval.FirstDerivative
			s.AzimuthalSecondDerivative[r][u] = interpolationInterval.SecondDerivative
		}
	}
	for r := range s.RadialPoints {
		for phi := range s.AzimuthalPoints {
			values := make([]float64, len(s.AngularPoints))
			for u := range s.Values[r] {
				values[u] = s.Values[r][u][phi]
			}
			interpolationInterval := InterpolationInterval{
				Start:   s.StartAngular,
				End:     s.EndAngular,
				Points:  s.AngularPoints,
				Weights: s.AngularWeights,
				Values:  values,
			}
			interpolationInterval.CalculateInterpolatingPolynomialCoefficientsAndFunctionDerivatives()
			for u := range s.AngularPoints {
				s.AngularFirstDerivative[r][u][phi] = interpolationInterval.FirstDerivative[u]
				s.AngularSecondDerivative[r][u][phi] = interpolationInterval.SecondDerivative[u]
			}
		}
	}
	for u := range s.AngularPoints {
		for phi := range s.AzimuthalPoints {
			values := make([]float64, len(s.RadialPoints))
			for r := range s.RadialPoints {
				values[r] = s.Values[r][u][phi]
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
				s.RadialFirstDerivative[r][u][phi] = interpolationInterval.FirstDerivative[r]
				s.RadialSecondDerivative[r][u][phi] = interpolationInterval.SecondDerivative[r]
			}
		}
	}
}

func (interval *IntegralVolume) CalculateHamiltonian(fn SphericalHamiltonian3d) {
	interval.calculateDerivatives()
	interval.KineticEnergy = make([][][]float64, len(interval.RadialPoints))
	interval.Potential = make([][][]float64, len(interval.RadialPoints))
	for r := range interval.RadialPoints {
		interval.KineticEnergy[r] = make([][]float64, len(interval.AngularPoints))
		interval.Potential[r] = make([][]float64, len(interval.AngularPoints))
		for u := range interval.AngularPoints {
			interval.KineticEnergy[r][u] = make([]float64, len(interval.AzimuthalPoints))
			interval.Potential[r][u] = make([]float64, len(interval.AzimuthalPoints))
			for phi := range interval.AzimuthalPoints {
				interval.KineticEnergy[r][u][phi], interval.Potential[r][u][phi] = fn(*interval, r, u, phi)
			}
		}
	}
}

func (interval *IntegralVolume) Integrate() {
	interval.IntegralKinetic = 0
	interval.IntegralPotential = 0
	for i := range interval.RadialWeights {
		r := interval.RadialPoints[i]
		for j := range interval.AngularWeights {
			th := interval.AngularPoints[j]
			for k := range interval.AzimuthalWeights {
				integratedKinetic := r * r * math.Sin(th) * interval.Values[i][j][k] * interval.KineticEnergy[i][j][k]
				integratedPotential := r * r * math.Sin(th) * interval.Values[i][j][k] * interval.Potential[i][j][k]
				interval.IntegralKinetic += integratedKinetic * interval.RadialWeights[i] * interval.AngularWeights[j] * interval.AzimuthalWeights[k]
				interval.IntegralPotential += integratedPotential * interval.RadialWeights[i] * interval.AngularWeights[j] * interval.AzimuthalWeights[k]
			}
		}
	}
}

func ParallelCalculations[T Interval](grid []T, fn func(T) error, threadLimit int) {
	var g errgroup.Group
	g.SetLimit(threadLimit)
	for i := range grid {
		interval := grid[i]
		g.Go(func() error {
			return fn(interval)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
