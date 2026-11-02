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
	calculateHydrogenPDOrbitalEnergy(2, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	calculateHydrogenPDOrbitalEnergy(3, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	calculateHydrogenPDOrbitalEnergy(3, 2, 2.0, 100, 24, 40, 1, 24, 40, 1)
	calculateHydrogenPDOrbitalEnergy(4, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	calculateHydrogenPDOrbitalEnergy(4, 2, 2.0, 100, 24, 40, 1, 24, 40, 1)
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
			MainQuantumNumber:   mainQuantumNumber,
			NumberOfIntervals:   numberOfIntervals,
			IntervalSize:        intervalSize,
			NumberOfGaussPoints: i,
		}
		//fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
		hyd.GenerateGrid()
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
	fmt.Printf("Calculation parameters \n Orbital %ds\n interval size: %f, number of intervals: %d, G-L points range: %d-%d\n", mainQuantumNumber, intervalSize, numberOfIntervals, gLpointsMin, gLpointsMax)
	fmt.Printf("Most accurate energy:  %.17g - calculated with %d G-L points\n", baseEnergy, optimalGLPointsNumber)
	fmt.Printf("Average energy: %.17g \n", average)
}

func calculateHydrogenPDOrbitalEnergy(
	mainQuantumNumber hydrogen.MainQuantumNumber,
	azimuthalQuantumNumber hydrogen.AzimuthalQuantumNumber,
	areaRadialDimension float64,
	numberOfAreas int,
	radialGLPointsMin int,
	radialGLPointsMax int,
	radialGLPointsStep int,
	angGLPointsMin int,
	angGLPointsMax int,
	angGLPointsStep int) {
	exactValue := 0.0
	switch mainQuantumNumber {
	case 2:
		exactValue = -0.125
	case 3:
		exactValue = -1.0 / 18.0
	case 4:
		exactValue = -1.0 / 32.0
	default:
		panic(mainQuantumNumber)
	}

	average := 0.0
	count := 0
	optimalRadialGLPointsNumber := 0
	optimalAngularGLPointsNumber := 0
	baseEnergy := 0.0
	delta := math.Abs(exactValue - baseEnergy)
	for rd := radialGLPointsMin; rd <= radialGLPointsMax; rd += radialGLPointsStep {
		for an := angGLPointsMin; an <= angGLPointsMax; an += angGLPointsStep {
			hyd := hydrogen.OrbitalPAtomic{
				MainQuantumNumber:       mainQuantumNumber,
				AzimuthalQuantumNumber:  azimuthalQuantumNumber,
				NumberOfAreas:           numberOfAreas,
				RadialNumberOfGLPoints:  rd,
				AngularNumberOfGLPoints: an,
				RadialIntervalSize:      areaRadialDimension,
			}

			//fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
			hyd.GenerateGrid()
			hyd.CalculateEnergy()
			//fmt.Printf("Calculated hydrogen ground state energy in atomic units is %.17g \n", hyd.Energy)
			newDelta := math.Abs(exactValue - hyd.Energy)
			if newDelta < delta {
				baseEnergy = hyd.Energy
				delta = newDelta
				optimalRadialGLPointsNumber = hyd.RadialNumberOfGLPoints
				optimalAngularGLPointsNumber = hyd.AngularNumberOfGLPoints
			}
			count++
			average += hyd.Energy
		}
	}
	average = average / float64(count)
	fmt.Printf("Calculation parameters \n Orbital n=%d l=%d\n Area radial size: %f, number of areas: %d, radial G-L points range: %d-%d\n angular G-L points range %d-%d\n",
		mainQuantumNumber, azimuthalQuantumNumber, areaRadialDimension, numberOfAreas, radialGLPointsMin, radialGLPointsMax, angGLPointsMin, angGLPointsMax)
	fmt.Printf("Most accurate energy:  %.17g - calculated with %d radial and %d angular G-L points\n", baseEnergy, optimalRadialGLPointsNumber, optimalAngularGLPointsNumber)
	fmt.Printf("Average energy: %.17g \n", average)
}
