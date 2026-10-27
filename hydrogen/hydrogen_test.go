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
		NumberOfAreas:           100,
		RadialNumberOfGLPoints:  32,
		AngularNumberOfGLPoints: 32,
		RadialIntervalSize:      2.0,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.RadialIntervalSize, hyd.NumberOfAreas, hyd.RadialNumberOfGLPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g\n", hyd.Energy)
}
