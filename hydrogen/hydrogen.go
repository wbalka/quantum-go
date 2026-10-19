package hydrogen

import (
	"errors"
	"math"
	"quantum-go/maths"
)

type RadialFunction interface {
	Evaluate(r float64) float64
}

type OrbitalS struct {
}

type MainQuantumNumber int

type OrbitalSAtomic struct {
	NumberOfIntervals   int
	IntervalSize        float64
	NumberOfGaussPoints int
	Grid                []maths.IntegralInterval
	Energy              float64
}

func (s OrbitalS) Evaluate(r float64) float64 {
	bohr := 0.52917721067                            // Bohr radius in angstroms
	r = r / bohr                                     // Convert r to Bohr units
	alpha := 1.0 / math.Sqrt(math.Pi*bohr*bohr*bohr) // Normalization constant for 1s orbital
	return alpha * math.Exp(-r)
}

func (s OrbitalSAtomic) Evaluate(r float64, mainQuantumNumber MainQuantumNumber) float64 {
	switch mainQuantumNumber {
	case 1:
		alpha := 1.0 / math.Sqrt(math.Pi) // Normalization constant for 1s orbital
		return alpha * math.Exp(-r)
	case 2:
		alpha := 1.0 / (4 * math.Sqrt(2*math.Pi)) // Normalization constant for 1s orbital
		return alpha * (2 - r) * math.Exp(-r/2)
	}
	return 0
}

func (s OrbitalSAtomic) Evaluate2s(r float64) float64 { // Convert r to Bohr units
	alpha := 1.0 / (4 * math.Sqrt(2*math.Pi)) // Normalization constant for 1s orbital
	return alpha * (2 - r) * math.Exp(-r/2)
}

// RadialDensity calculates probability density in a sphere of r radius
func (s OrbitalS) RadialDensity(r float64) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s OrbitalSAtomic) RadialDensity(r float64, mainQuantumNumber MainQuantumNumber) float64 {
	psi := s.Evaluate(r, mainQuantumNumber)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s *OrbitalSAtomic) GenerateGrid(mainQuantumNumber MainQuantumNumber) error {
	if mainQuantumNumber < 1 || mainQuantumNumber > 2 {
		return errors.New("Invalid main quantum number")
	}
	s.Grid = make([]maths.IntegralInterval, s.NumberOfIntervals)
	for i := range s.NumberOfIntervals {
		start := float64(i) * s.IntervalSize
		end := start + s.IntervalSize
		points, weights, _ := maths.GaussLegendrePointsWeights(s.NumberOfGaussPoints)
		values := make([]float64, s.NumberOfGaussPoints)
		for j := range points {
			points[j] = (end-start)/2*points[j] + (end+start)/2
			weights[j] = (end - start) / 2 * weights[j]
			values[j] = s.Evaluate(points[j], mainQuantumNumber)
		}
		s.Grid[i] = maths.IntegralInterval{
			Start:   start,
			End:     end,
			Points:  points,
			Weights: weights,
			Values:  values,
		}
	}
	return nil
}

func Hydrogen1SHamiltonian(interval maths.IntegralInterval, i int) float64 {
	return -0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
}

func (s *OrbitalSAtomic) CalculateEnergy() {
	for i := range s.Grid {
		interval := &s.Grid[i]
		interval.IntegralValue = 0
		interval.CalculateHamiltonian(Hydrogen1SHamiltonian)
		interval.Integrate()
		s.Energy += interval.IntegralValue
	}
	s.Energy *= 4 * math.Pi
}
