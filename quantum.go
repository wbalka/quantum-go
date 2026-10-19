package main

import (
	"fmt"
	"math"
	"quantum-go/hydrogen"
)

func main() {
	calculateHydrogenSOrbitalEnergy(1, 2.0, 100, 14, 80)
	calculateHydrogenSOrbitalEnergy(2, 2.0, 100, 14, 80)
	calculateHydrogenSOrbitalEnergy(2, 1.0, 200, 14, 80)
	/*
		 	calculateHydrogenSOrbitalEnergy(2, 2.0, 1000, 8, 100)
			calculateHydrogenSOrbitalEnergy(2, 1.0, 2000, 8, 100)
	*/
}

func calculateHydrogenSOrbitalEnergy(mainQuantumNumber hydrogen.MainQuantumNumber, intervalSize float64, numberOfIntervals int, gLpointsMin int, gLpointsMax int) {
	exactValue := 0.0
	switch mainQuantumNumber {
	case 1:
		exactValue = -0.5
	case 2:
		exactValue = -0.125
	default:
		panic(mainQuantumNumber)
	}
	average := 0.0
	count := 0
	optimalGLPointsNumber := 0
	baseEnergy := 0.0
	delta := math.Abs(exactValue - baseEnergy)
	for i := gLpointsMin; i <= gLpointsMax; i++ {
		hyd := hydrogen.OrbitalSAtomic{
			NumberOfIntervals:   numberOfIntervals,
			IntervalSize:        intervalSize,
			NumberOfGaussPoints: i,
		}
		//fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
		hyd.GenerateGrid(mainQuantumNumber)
		hyd.CalculateEnergy()
		//fmt.Printf("Calculated hydrogen ground state energy in atomic units is %.17g \n", hyd.Energy)
		newDelta := math.Abs(exactValue - hyd.Energy)
		if newDelta < delta {
			baseEnergy = hyd.Energy
			delta = newDelta
			optimalGLPointsNumber = hyd.NumberOfGaussPoints
		}
		count++
		average += hyd.Energy
	}
	average = average / float64(count)
	fmt.Printf("Calculation parameters \n interval size: %f, number of intervals: %d, G-L points range: %d-%d\n", intervalSize, numberOfIntervals, gLpointsMin, gLpointsMax)
	fmt.Printf("Best %ds energy:  %.17g - calculated with %d G-L points\n", mainQuantumNumber, baseEnergy, optimalGLPointsNumber)
	fmt.Printf("Average energy: %.17g \n", average)
}
