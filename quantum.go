package main

import (
	"fmt"
	"quantum-go/hydrogen"
)

func main() {
	hyd := hydrogen.OrbitalSAtomic{
		NumberOfIntervals:   10000,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 32,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Calculated hydrogen ground state energy in atomic units is %.17g", hyd.Energy)
}
