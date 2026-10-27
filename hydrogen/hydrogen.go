package hydrogen

import (
	"errors"
	"fmt"
	"math"
	"quantum-go/maths"
)

type MainQuantumNumber int

type RadialFunction interface {
	Evaluate(r float64) float64
}

type OrbitalS struct {
}

type OrbitalSAtomic struct {
	MainQuantumNumber
	NumberOfIntervals   int
	IntervalSize        float64
	NumberOfGaussPoints int
	Grid                []maths.RadialInterval
	Energy              float64
}

type OrbitalPAtomic struct {
	MainQuantumNumber
	NumberOfAreas           int
	RadialNumberOfGLPoints  int
	AngularNumberOfGLPoints int
	RadialIntervalSize      float64
	Grid                    []maths.IntegralArea
	Energy                  float64
	Kinetic                 float64
	Potential               float64
}

func (s OrbitalS) Evaluate(r float64) float64 {
	bohr := 0.52917721067                            // Bohr radius in angstroms
	r = r / bohr                                     // Convert r to Bohr units
	alpha := 1.0 / math.Sqrt(math.Pi*bohr*bohr*bohr) // Normalization constant for 1s orbital
	return alpha * math.Exp(-r)
}

func (s OrbitalSAtomic) Evaluate(r float64) float64 {
	switch s.MainQuantumNumber {
	case 1:
		alpha := 1.0 / math.Sqrt(math.Pi) // Normalization constant for 1s orbital
		return alpha * math.Exp(-r)
	case 2:
		alpha := 1.0 / (4 * math.Sqrt(2*math.Pi)) // Normalization constant for 1s orbital
		return alpha * (2 - r) * math.Exp(-r/2)
	}
	return 0
}

func (s OrbitalPAtomic) Evaluate(r float64, u float64) float64 {
	switch s.MainQuantumNumber {
	case 2:
		alpha := 1.0 / (4 * math.Sqrt(2*math.Pi))
		return alpha * r * math.Exp(-r/2) * u
	default:
		return 0
	}
}

// RadialDensity calculates probability density in a sphere of r radius
func (s OrbitalS) RadialDensity(r float64) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s OrbitalSAtomic) RadialDensity(r float64, mainQuantumNumber MainQuantumNumber) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s *OrbitalSAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 1 || s.MainQuantumNumber > 2 {
		return errors.New("Invalid main quantum number")
	}
	s.Grid = make([]maths.RadialInterval, s.NumberOfIntervals)
	for i := range s.NumberOfIntervals {
		start := float64(i) * s.IntervalSize
		end := start + s.IntervalSize
		points, weights, _ := maths.GaussLegendrePointsWeights(s.NumberOfGaussPoints)
		values := make([]float64, s.NumberOfGaussPoints)
		for j := range points {
			points[j] = (end-start)/2*points[j] + (end+start)/2
			weights[j] = (end - start) / 2 * weights[j]
			values[j] = s.Evaluate(points[j])
		}
		s.Grid[i] = maths.RadialInterval{
			Start:   start,
			End:     end,
			Points:  points,
			Weights: weights,
			Values:  values,
		}
	}
	return nil
}

func (s *OrbitalPAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 2 || s.MainQuantumNumber > 2 {
		return errors.New("Invalid main quantum number")
	}
	s.Grid = make([]maths.IntegralArea, s.NumberOfAreas)
	for i := range s.NumberOfAreas {
		start := float64(i) * s.RadialIntervalSize
		end := start + s.RadialIntervalSize
		radialPoints, radialWeights, _ := maths.GaussLegendrePointsWeights(s.RadialNumberOfGLPoints)
		angularPoints, angularWeights, _ := maths.GaussLegendrePointsWeights(s.AngularNumberOfGLPoints)
		values := make([][]float64, s.RadialNumberOfGLPoints)
		for j := range radialPoints {
			radialPoints[j] = (end-start)/2*radialPoints[j] + (end+start)/2
			radialWeights[j] = (end - start) / 2 * radialWeights[j]
			values[j] = make([]float64, s.AngularNumberOfGLPoints)
			for k := range angularPoints {
				values[j][k] = s.Evaluate(radialPoints[j], angularPoints[k])
			}
		}
		s.Grid[i] = maths.IntegralArea{
			StartRadial:         start,
			EndRadial:           end,
			StartAngular:        -1,
			EndAngular:          1,
			RadialIntervalSize:  s.RadialIntervalSize,
			AngularIntervalSize: 2.0,
			RadialPoints:        radialPoints,
			RadialWeights:       radialWeights,
			AngularPoints:       angularPoints,
			AngularWeights:      angularWeights,
			Values:              values,
		}
	}
	return nil
}

func HydrogenSHamiltonian(interval maths.RadialInterval, i int) float64 {
	return -0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
}

func (s *OrbitalSAtomic) CalculateEnergy() {
	for i := range s.Grid {
		interval := &s.Grid[i]
		interval.IntegralValue = 0
		interval.CalculateHamiltonian(HydrogenSHamiltonian)
		interval.Integrate()
		s.Energy += interval.IntegralValue
	}
	s.Energy *= 4 * math.Pi
}

func HydrogenPHamiltonian(interval maths.IntegralArea, r int, u int) (float64, float64) {
	kineticRadial := interval.RadialSecondDerivative[r][u] + 2*interval.RadialFirstDerivative[r][u]/interval.RadialPoints[r]
	kineticAngular := (-2*interval.AngularPoints[u]*interval.AngularFirstDerivative[r][u] + (1-interval.AngularPoints[u]*interval.AngularPoints[u])*interval.AngularSecondDerivative[r][u]) / (interval.RadialPoints[r] * interval.RadialPoints[r])
	potential := interval.Values[r][u] / interval.RadialPoints[r]
	return -0.5 * (kineticRadial + kineticAngular), -potential
}

func (s *OrbitalPAtomic) CalculateEnergy() {
	for i := range s.Grid {
		interval := &s.Grid[i]
		interval.IntegralKinetic = 0
		interval.IntegralPotential = 0
		interval.CalculateHamiltonian(HydrogenPHamiltonian)

	}
	for i := range s.Grid {
		interval := &s.Grid[i]
		interval.Integrate()
		s.Kinetic += interval.IntegralKinetic
		s.Potential += interval.IntegralPotential
	}
	s.Kinetic *= 2 * math.Pi
	s.Potential *= 2 * math.Pi
	s.Energy = s.Kinetic + s.Potential
	fmt.Printf("Kinetic %0.17g potential %0.17g \n", s.Kinetic, s.Potential)
}
