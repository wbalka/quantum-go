package main

import (
	"fmt"
	"quantum-go/hydrogen"
)

func main() {
	for i := 3; i <= 100; i++ {
		hyd := hydrogen.OrbitalSAtomic{
			NumberOfIntervals:   100,
			IntervalSize:        2.0,
			NumberOfGaussPoints: i,
		}
		fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
		hyd.GenerateGrid()
		hyd.CalculateEnergy()
		fmt.Printf("Calculated hydrogen ground state energy in atomic units is %.17g \n", hyd.Energy)
	}
}
