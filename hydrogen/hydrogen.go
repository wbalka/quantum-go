package hydrogen

import (
	"math"
	"quantum-go/maths"
)

type RadialFunction interface {
	Evaluate(r float64) float64
}

type OrbitalS struct {
}

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

func (s OrbitalSAtomic) Evaluate(r float64) float64 { // Convert r to Bohr units
	alpha := 1.0 / math.Sqrt(math.Pi) // Normalization constant for 1s orbital
	return alpha * math.Exp(-r)
}

// RadialDensity oblicza gęstość prawdopodobieństwa znalezienia elektronu w sferze o promieniu r
func (s OrbitalS) RadialDensity(r float64) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s OrbitalSAtomic) RadialDensity(r float64) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s *OrbitalSAtomic) GenerateGrid() {
	s.Grid = make([]maths.IntegralInterval, s.NumberOfIntervals)
	for i := range s.NumberOfIntervals {
		start := float64(i) * s.IntervalSize
		end := start + s.IntervalSize
		points, weights, _ := maths.GaussLegendrePointsWeights(s.NumberOfGaussPoints)
		values := make([]float64, s.NumberOfGaussPoints)
		//values[0] = s.Evaluate(start)
		//values[s.numberOfGaussPoints+1] = s.Evaluate(end)
		for j := range points {
			points[j] = (end-start)/2*points[j] + (end+start)/2
			weights[j] = (end - start) / 2 * weights[j]
			values[j] = s.Evaluate(points[j])
		}
		s.Grid[i] = maths.IntegralInterval{
			Start:   start,
			End:     end,
			Points:  points,
			Weights: weights,
			Values:  values,
		}
	}
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
