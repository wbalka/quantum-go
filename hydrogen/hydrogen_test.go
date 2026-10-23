package hydrogen

import (
	"fmt"
	"testing"
)

func TestHydrogen1sEnergy(t *testing.T) {
	hyd := OrbitalSAtomic{
		MainQuantumNumber:   1,
		NumberOfIntervals:   100,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 32,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g", hyd.Energy)
}

func TestHydrogen2sEnergy(t *testing.T) {
	hyd := OrbitalSAtomic{
		MainQuantumNumber:   2,
		NumberOfIntervals:   10000,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 32,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g", hyd.Energy)
}

func TestHydrogen2pEnergy(t *testing.T) {
	hyd := OrbitalPAtomic{
		MainQuantumNumber:       2,
		numberOfAreas:           100,
		radialNumberOfGLPoints:  32,
		angularNumberOfGLPoints: 32,
		radialIntervalSize:      2.0,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.radialIntervalSize, hyd.numberOfAreas, hyd.radialNumberOfGLPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g", hyd.Energy)
}
