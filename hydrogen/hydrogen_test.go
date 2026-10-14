package hydrogen

import (
	"fmt"
	"testing"
)

func TestHydrogenEnergy(t *testing.T) {
	hyd := OrbitalSAtomic{
		NumberOfIntervals:   100,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 32,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g", hyd.Energy)
}
