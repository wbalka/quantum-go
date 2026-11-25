package hydrogen

import (
	"fmt"
	"testing"
)

func TestHydrogen1sEnergy(t *testing.T) {
	hyd := SphericalOrbitalAtomic{
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
	hyd := SphericalOrbitalAtomic{
		MainQuantumNumber:   2,
		NumberOfIntervals:   10000,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 64,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.IntervalSize, hyd.NumberOfIntervals, hyd.NumberOfGaussPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g", hyd.Energy)
}

func TestHydrogen3sEnergy(t *testing.T) {
	hyd := SphericalOrbitalAtomic{
		MainQuantumNumber:   3,
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
	hyd := AxialOrbitalAtomic{
		MainQuantumNumber:       2,
		AzimuthalQuantumNumber:  1,
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

func TestHydrogen3pEnergy(t *testing.T) {
	hyd := AxialOrbitalAtomic{
		MainQuantumNumber:       3,
		AzimuthalQuantumNumber:  1,
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

func TestHydrogen3dEnergy(t *testing.T) {
	hyd := AxialOrbitalAtomic{
		MainQuantumNumber:       3,
		AzimuthalQuantumNumber:  2,
		NumberOfAreas:           10000,
		RadialNumberOfGLPoints:  32,
		AngularNumberOfGLPoints: 32,
		RadialIntervalSize:      2.0,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of intervals %d, number of Gauss points %d \n", hyd.RadialIntervalSize, hyd.NumberOfAreas, hyd.RadialNumberOfGLPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g\n", hyd.Energy)
}

func TestHydrogenEnergyMagnetic4F(t *testing.T) {
	for m := -3; m <= 3; m++ {
		hyd := Universal3DOrbitalAtomic{
			MainQuantumNumber:          4,
			AzimuthalQuantumNumber:     3,
			MagneticQuantumNumber:      MagneticQuantumNumber(m),
			NumberOfLayers:             100,
			NumberOfAzimuthalIntervals: 4,
			RadialNumberOfGLPoints:     16,
			AngularNumberOfGLPoints:    32,
			AzimuthalNumberOfGLPoints:  16,
			RadialIntervalSize:         2.0,
		}
		fmt.Printf("Calculation parameters: interval size %f, number of layers %d, number of Gauss points %d \n", hyd.RadialIntervalSize, hyd.NumberOfLayers, hyd.RadialNumberOfGLPoints)
		fmt.Printf("n=%d l=%d m=%d\n", hyd.MainQuantumNumber, hyd.AzimuthalQuantumNumber, hyd.MagneticQuantumNumber)
		hyd.GenerateGrid()
		hyd.CalculateEnergy()
		fmt.Printf("Eneergy is %.17g\n", hyd.Energy)
	}
}

func TestHydrogen3dEnergyMagnetic5F(t *testing.T) {
	for m := -3; m <= 3; m++ {
		hyd := Universal3DOrbitalAtomic{
			MainQuantumNumber:          5,
			AzimuthalQuantumNumber:     3,
			MagneticQuantumNumber:      MagneticQuantumNumber(m),
			NumberOfLayers:             100,
			NumberOfAzimuthalIntervals: 8,
			RadialNumberOfGLPoints:     16,
			AngularNumberOfGLPoints:    32,
			AzimuthalNumberOfGLPoints:  16,
			RadialIntervalSize:         2.0,
		}
		fmt.Printf("Calculation parameters: interval size %f, number of layers %d, number of Gauss points %d \n", hyd.RadialIntervalSize, hyd.NumberOfLayers, hyd.RadialNumberOfGLPoints)
		fmt.Printf("n=%d l=%d m=%d\n", hyd.MainQuantumNumber, hyd.AzimuthalQuantumNumber, hyd.MagneticQuantumNumber)
		hyd.GenerateGrid()
		hyd.CalculateEnergy()
		fmt.Printf("Eneergy is %.17g\n", hyd.Energy)
	}
}

func TestGenerateGrid3d(t *testing.T) {
	hyd := Universal3DOrbitalAtomic{
		MainQuantumNumber:          2,
		AzimuthalQuantumNumber:     0,
		MagneticQuantumNumber:      0,
		NumberOfLayers:             100,
		NumberOfAzimuthalIntervals: 4,
		RadialNumberOfGLPoints:     16,
		AngularNumberOfGLPoints:    16,
		AzimuthalNumberOfGLPoints:  16,
		RadialIntervalSize:         2.0,
	}
	hyd.GenerateGrid()
	value_1 := hyd.Grid[0].Values[10][7][11]
	value_3 := hyd.Grid[100].Values[10][7][11]
	hyd_2 := SphericalOrbitalAtomic{
		MainQuantumNumber:   2,
		NumberOfIntervals:   100,
		IntervalSize:        2.0,
		NumberOfGaussPoints: 16,
	}
	hyd_2.GenerateGrid()
	value_2 := hyd_2.Grid[0].Values[10]
	fmt.Printf("Received value %0.17g, expected %0.17g\n", value_1, value_2)
	fmt.Println(value_3)
	if value_1 != value_2 {
		t.Error("Bad grid generator!")
	}
}

func TestGenerateGrid3dMagnetic(t *testing.T) {
	hyd := Universal3DOrbitalAtomic{
		MainQuantumNumber:          3,
		AzimuthalQuantumNumber:     2,
		MagneticQuantumNumber:      -2,
		NumberOfLayers:             100,
		NumberOfAzimuthalIntervals: 4,
		RadialNumberOfGLPoints:     16,
		AngularNumberOfGLPoints:    16,
		AzimuthalNumberOfGLPoints:  16,
		RadialIntervalSize:         2.0,
	}
	hyd.GenerateGrid()
	value_1 := hyd.Grid[2].Values[10][7][11]
	value_2 := hyd.Grid[102].Values[10][7][11]
	fmt.Printf("Received value %0.17g, expected %0.17g\n", value_2, value_1)
	if value_1 == value_2 {
		t.Error("Bad grid generator!")
	}
}
