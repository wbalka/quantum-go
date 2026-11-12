package hydrogen

import (
	"errors"
	"fmt"
	"math"
	"quantum-go/maths"
	"runtime"

	"golang.org/x/sync/errgroup"
)

type MainQuantumNumber int

type AzimuthalQuantumNumber int

type SphericalOrbitalAtomic struct {
	MainQuantumNumber
	maths.WaveFunction
	NumberOfIntervals   int
	IntervalSize        float64
	NumberOfGaussPoints int
	Grid                []maths.RadialInterval
	Energy              float64
}

type AxialOrbitalAtomic struct {
	MainQuantumNumber
	AzimuthalQuantumNumber
	maths.WaveFunction
	NumberOfAreas           int
	RadialNumberOfGLPoints  int
	AngularNumberOfGLPoints int
	RadialIntervalSize      float64
	Grid                    []maths.IntegralArea
	Energy                  float64
	Kinetic                 float64
	Potential               float64
}

func (s SphericalOrbitalAtomic) Evaluate(r float64) float64 {
	return s.WaveFunction(r, 0)
}

func (s AxialOrbitalAtomic) Evaluate(r float64, u float64) float64 {
	return s.WaveFunction(r, u)
}

func (s SphericalOrbitalAtomic) RadialDensity(r float64, mainQuantumNumber MainQuantumNumber) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s *SphericalOrbitalAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 1 || s.MainQuantumNumber > 5 {
		return errors.New("Invalid main quantum number")
	}
	s.WaveFunction, _ = maths.GetWaveFunction(int(s.MainQuantumNumber), 0, 0, 1)
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

func (s *AxialOrbitalAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 2 || s.MainQuantumNumber > 5 {
		return errors.New("Invalid main quantum number")
	}
	s.WaveFunction, _ = maths.GetWaveFunction(int(s.MainQuantumNumber), int(s.AzimuthalQuantumNumber), 0, 1)
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

/*
Hamiltonian of hydrogen atom in atomic units (a.u.) for spherical symmetry.

Operator:
H^ = -1/2 * [ d2/dr2 + 2/r * d/dr ] - 1/r

Where:

	r - radial distance from nucleus
*/
func HydrogenSHamiltonian(interval maths.RadialInterval, i int) float64 {
	return -0.5*(interval.SecondDerivative[i]+2*interval.FirstDerivative[i]/interval.Points[i]) - interval.Values[i]/interval.Points[i]
}

func (s *SphericalOrbitalAtomic) CalculateEnergy() {
	var g errgroup.Group
	g.SetLimit(runtime.NumCPU())
	for i := range s.Grid {
		interval := &s.Grid[i]
		g.Go(func() error {
			return calculateRadialInterval(interval)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	for _, interval := range s.Grid {
		s.Energy += interval.IntegralValue
	}
	s.Energy *= 4 * math.Pi
}

func calculateRadialInterval(interval *maths.RadialInterval) error {
	interval.IntegralValue = 0
	interval.CalculateHamiltonian(HydrogenSHamiltonian)
	interval.Integrate()
	return nil
}

/*
Hamiltonian of hydrogen atom in atomic units (a.u.) for axial symmetry.
Substituion used: u = cos(theta), assuming d/d(phi) = 0.

Operator:
H^ = -1/2 * [ 1/r^2 * d/dr(r^2 * d/dr) + 1/r^2 * d/du((1 - u^2) * d/du) ] - 1/r

Where:

	r - radial distance from nucleus
	u - cosine of polar angle (cos(theta)), lies within range [-1, 1]
*/
func HydrogenPHamiltonian(interval maths.IntegralArea, r int, u int) (float64, float64) {
	kineticRadial := interval.RadialSecondDerivative[r][u] + 2*interval.RadialFirstDerivative[r][u]/interval.RadialPoints[r]
	kineticAngular := (-2*interval.AngularPoints[u]*interval.AngularFirstDerivative[r][u] + (1-interval.AngularPoints[u]*interval.AngularPoints[u])*interval.AngularSecondDerivative[r][u]) / (interval.RadialPoints[r] * interval.RadialPoints[r])
	potential := interval.Values[r][u] / interval.RadialPoints[r]
	return -0.5 * (kineticRadial + kineticAngular), -potential
}

func (s *AxialOrbitalAtomic) CalculateEnergy() {
	var g errgroup.Group
	g.SetLimit(runtime.NumCPU())
	for i := range s.Grid {
		interval := &s.Grid[i]
		g.Go(func() error {
			return calculateAxialArea(interval)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	for _, interval := range s.Grid {
		s.Kinetic += interval.IntegralKinetic
		s.Potential += interval.IntegralPotential
	}
	s.Kinetic *= 2 * math.Pi
	s.Potential *= 2 * math.Pi
	s.Energy = s.Kinetic + s.Potential
	//fmt.Printf("Kinetic %0.17g potential %0.17g \n", s.Kinetic, s.Potential)
}

func calculateAxialArea(area *maths.IntegralArea) error {
	area.IntegralKinetic = 0
	area.IntegralPotential = 0
	area.CalculateHamiltonian(HydrogenPHamiltonian)
	area.Integrate()
	return nil
}
