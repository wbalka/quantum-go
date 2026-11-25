package hydrogen

import (
	"errors"
	"fmt"
	"math"
	"quantum-go/maths"
	"runtime"
)

type MainQuantumNumber int

type AzimuthalQuantumNumber int

type MagneticQuantumNumber int

type SphericalOrbitalAtomic struct {
	MainQuantumNumber
	maths.WaveFunction
	NumberOfIntervals   int
	IntervalSize        float64
	NumberOfGaussPoints int
	Grid                []*maths.RadialInterval
	Energy              float64
	HamiltonianFn       maths.RadialHamiltonian
}

type AxialOrbitalAtomic struct {
	MainQuantumNumber
	AzimuthalQuantumNumber
	maths.WaveFunction
	NumberOfAreas           int
	RadialNumberOfGLPoints  int
	AngularNumberOfGLPoints int
	RadialIntervalSize      float64
	Grid                    []*maths.IntegralArea
	Energy                  float64
	Kinetic                 float64
	Potential               float64
	HamiltonianFn           maths.SphericalHamiltonian2d
}

type Universal3DOrbitalAtomic struct {
	MainQuantumNumber
	AzimuthalQuantumNumber
	MagneticQuantumNumber
	maths.WaveFunction
	NumberOfLayers             int
	NumberOfAzimuthalIntervals int
	RadialNumberOfGLPoints     int
	AngularNumberOfGLPoints    int
	AzimuthalNumberOfGLPoints  int
	RadialIntervalSize         float64
	AzimuthalIntervalSize      float64
	Grid                       []*maths.IntegralVolume
	Energy                     float64
	Kinetic                    float64
	Potential                  float64
	HamiltonianFn              maths.SphericalHamiltonian3d
}

func (s SphericalOrbitalAtomic) Evaluate(r float64) float64 {
	return s.WaveFunction(r, 0, 0)
}

func (s AxialOrbitalAtomic) Evaluate(r float64, u float64) float64 {
	return s.WaveFunction(r, u, 0)
}

func (s Universal3DOrbitalAtomic) Evaluate(r float64, u float64, phi float64) float64 {
	return s.WaveFunction(r, u, phi)
}

func (s SphericalOrbitalAtomic) RadialDensity(r float64, mainQuantumNumber MainQuantumNumber) float64 {
	psi := s.Evaluate(r)
	return 4.0 * math.Pi * r * r * psi * psi
}

func (s *SphericalOrbitalAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 1 || s.MainQuantumNumber > 5 {
		return errors.New("Invalid main quantum number")
	}
	s.HamiltonianFn = HydrogenSHamiltonian
	s.WaveFunction, _ = maths.GetWaveFunction(int(s.MainQuantumNumber), 0, 0, 1)
	s.Grid = make([]*maths.RadialInterval, s.NumberOfIntervals)
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
		s.Grid[i] = &maths.RadialInterval{
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
	s.HamiltonianFn = HydrogenPHamiltonian
	s.WaveFunction, _ = maths.GetWaveFunction(int(s.MainQuantumNumber), int(s.AzimuthalQuantumNumber), 0, 1)
	s.Grid = make([]*maths.IntegralArea, s.NumberOfAreas)
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
		s.Grid[i] = &maths.IntegralArea{
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

func (s *Universal3DOrbitalAtomic) GenerateGrid() error {
	if s.MainQuantumNumber < 2 || s.MainQuantumNumber > 5 {
		return errors.New("Invalid main quantum number")
	} else if s.AzimuthalQuantumNumber < 0 || s.AzimuthalQuantumNumber > AzimuthalQuantumNumber(s.MainQuantumNumber-1) {
		return errors.New("Invalid azimuthal quantum number")
	} else if s.MagneticQuantumNumber < -MagneticQuantumNumber(s.AzimuthalQuantumNumber) || s.MagneticQuantumNumber > MagneticQuantumNumber(s.AzimuthalQuantumNumber) {
		return errors.New("Invalid magnetic quantum number")
	}
	s.HamiltonianFn = HydrogenDHamiltonian
	s.WaveFunction, _ = maths.GetWaveFunction(int(s.MainQuantumNumber), int(s.AzimuthalQuantumNumber), int(s.MagneticQuantumNumber), 1)
	s.Grid = make([]*maths.IntegralVolume, s.NumberOfAzimuthalIntervals*s.NumberOfLayers)
	gridIndex := 0
	s.AzimuthalIntervalSize = 2 * math.Pi / float64(s.NumberOfAzimuthalIntervals)
	baseRadialPoints, baseRadialWeights, _ := maths.GaussLegendrePointsWeights(s.RadialNumberOfGLPoints)
	baseAzimuthalPoints, baseAzimuthalWeights, _ := maths.GaussLegendrePointsWeights(s.AzimuthalNumberOfGLPoints)
	angularPoints, angularWeights, _ := maths.GaussLegendrePointsWeights(s.AngularNumberOfGLPoints)
	for az := range s.NumberOfAzimuthalIntervals {
		azimuthalPoints := make([]float64, len(baseAzimuthalPoints))
		azimuthalWeights := make([]float64, len(baseAzimuthalWeights))
		azStart := float64(az) * s.AzimuthalIntervalSize
		azEnd := azStart + s.AzimuthalIntervalSize
		for k := range azimuthalPoints {
			azimuthalPoints[k] = (azEnd-azStart)/2*baseAzimuthalPoints[k] + (azEnd+azStart)/2
			azimuthalWeights[k] = (azEnd - azStart) / 2 * baseAzimuthalWeights[k]
		}
		for ly := range s.NumberOfLayers {
			rStart := float64(ly) * s.RadialIntervalSize
			rEnd := rStart + s.RadialIntervalSize
			radialPoints := make([]float64, len(baseRadialPoints))
			radialWeights := make([]float64, len(baseRadialWeights))
			for i := range baseRadialPoints {
				radialPoints[i] = (rEnd-rStart)/2*baseRadialPoints[i] + (rEnd+rStart)/2
				radialWeights[i] = (rEnd - rStart) / 2 * baseRadialWeights[i]
			}
			values := make([][][]float64, s.RadialNumberOfGLPoints)
			for i := range baseRadialPoints {
				values[i] = make([][]float64, s.AngularNumberOfGLPoints)
				for j := range angularPoints {
					values[i][j] = make([]float64, s.AzimuthalNumberOfGLPoints)
					for k := range azimuthalPoints {
						values[i][j][k] = s.Evaluate(radialPoints[i], angularPoints[j], azimuthalPoints[k])
					}
				}
			}
			s.Grid[gridIndex] = &maths.IntegralVolume{
				StartRadial:           rStart,
				EndRadial:             rEnd,
				StartAngular:          -1,
				EndAngular:            1,
				RadialIntervalSize:    s.RadialIntervalSize,
				AngularIntervalSize:   2.0,
				StartAzimuthal:        azStart,
				EndAzimuthal:          azEnd,
				AzimuthalIntervalSize: s.AzimuthalIntervalSize,
				RadialPoints:          radialPoints,
				RadialWeights:         radialWeights,
				AngularPoints:         angularPoints,
				AngularWeights:        angularWeights,
				AzimuthalPoints:       azimuthalPoints,
				AzimuthalWeights:      azimuthalWeights,
				Values:                values,
			}
			gridIndex++
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
	maths.ParallelCalculations(s.Grid,
		func(interval *maths.RadialInterval) error {
			interval.CalculateHamiltonian(s.HamiltonianFn)
			interval.Integrate()
			return nil
		},
		runtime.NumCPU())
	for _, interval := range s.Grid {
		s.Energy += interval.IntegralValue
	}
	s.Energy *= 4 * math.Pi
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

/*
Hamiltonian of hydrogen atom in atomic units (a.u.) without any assumed symmetry.
Substituion used: u = cos(theta).

Operator:
H^ = -1/2 * [ 1/r^2 * d/dr(r^2 * d/dr) + 1/r^2 * d/du((1 - u^2) * d/du) + 1/(r^2 * (1 - u^2)) * d2/dphi2] - 1/r

Where:

	r - radial distance from nucleus
	u - cosine of polar angle (cos(theta)), lies within range [-1, 1]
	phi - azimuthal angle [0, 2pi]
*/
func HydrogenDHamiltonian(interval maths.IntegralVolume, r int, u int, ph int) (float64, float64) {
	kineticRadial := interval.RadialSecondDerivative[r][u][ph] + 2*interval.RadialFirstDerivative[r][u][ph]/interval.RadialPoints[r]
	kineticAngular := (-2*interval.AngularPoints[u]*interval.AngularFirstDerivative[r][u][ph] + (1-interval.AngularPoints[u]*interval.AngularPoints[u])*interval.AngularSecondDerivative[r][u][ph]) / (interval.RadialPoints[r] * interval.RadialPoints[r])
	kineticAzimuthal := interval.AzimuthalSecondDerivative[r][u][ph] / (interval.RadialPoints[r] * interval.RadialPoints[r] * (1 - interval.AngularPoints[u]*interval.AngularPoints[u]))
	potential := interval.Values[r][u][ph] / interval.RadialPoints[r]
	return -0.5 * (kineticRadial + kineticAngular + kineticAzimuthal), -potential
}

func (s *AxialOrbitalAtomic) CalculateEnergy() {
	maths.ParallelCalculations(s.Grid,
		func(interval *maths.IntegralArea) error {
			interval.CalculateHamiltonian(s.HamiltonianFn)
			interval.Integrate()
			return nil
		},
		runtime.NumCPU())
	for _, interval := range s.Grid {
		s.Kinetic += interval.IntegralKinetic
		s.Potential += interval.IntegralPotential
	}
	s.Kinetic *= 2 * math.Pi
	s.Potential *= 2 * math.Pi
	s.Energy = s.Kinetic + s.Potential
	fmt.Printf("Kinetic %0.17g potential %0.17g \n", s.Kinetic, s.Potential)
}

func (s *Universal3DOrbitalAtomic) CalculateEnergy() {
	maths.ParallelCalculations(s.Grid,
		func(interval *maths.IntegralVolume) error {
			interval.CalculateHamiltonian(s.HamiltonianFn)
			interval.Integrate()
			return nil
		},
		runtime.NumCPU())
	for _, interval := range s.Grid {
		s.Kinetic += interval.IntegralKinetic
		s.Potential += interval.IntegralPotential
	}
	s.Energy = s.Kinetic + s.Potential
	fmt.Printf("Kinetic %0.17g potential %0.17g \n", s.Kinetic, s.Potential)
}

func (s *Universal3DOrbitalAtomic) CalculateEnergySeq() {
	for _, interval := range s.Grid {
		interval.CalculateHamiltonian(s.HamiltonianFn)
		interval.Integrate()
		s.Kinetic += interval.IntegralKinetic
		s.Potential += interval.IntegralPotential
	}
	s.Energy = s.Kinetic + s.Potential
	fmt.Printf("Kinetic %0.17g potential %0.17g \n", s.Kinetic, s.Potential)
}
