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

func TestHydrogen3dEnergyMagnetic(t *testing.T) {
	hyd := Universal3DOrbitalAtomic{
		MainQuantumNumber:          3,
		AzimuthalQuantumNumber:     2,
		MagneticQuantumNumber:      -1,
		NumberOfLayers:             100,
		NumberOfAzimuthalIntervals: 4,
		RadialNumberOfGLPoints:     16,
		AngularNumberOfGLPoints:    96,
		AzimuthalNumberOfGLPoints:  16,
		RadialIntervalSize:         2.0,
	}
	fmt.Printf("Calculation parameters: interval size %f, number of layers %d, number of Gauss points %d \n", hyd.RadialIntervalSize, hyd.NumberOfLayers, hyd.RadialNumberOfGLPoints)
	hyd.GenerateGrid()
	hyd.CalculateEnergy()
	fmt.Printf("Eneergy is %.17g\n", hyd.Energy)
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

/**
Świetny plan — powodzenia z przepisywaniem!
Pamiętaj tylko o trzech drobiazgach przy zmianie zmiennej z u na θ:
• Przeskalowanie węzłów i wag: z \([-1, 1]\) na \([0, \pi]\) (mnożnik przedziału to \(\frac{\pi }{2}\)).
• Jakobian: w elemencie objętości zamiast r² dr du dφ wchodzi \(r^2 \sin\theta \, dr \, d\theta \, d\phi\).
• Laplasjan w HydrogenDHamiltonian: część biegunowa to teraz \(\frac{1}{r^2 \sin\theta}\frac{\partial}{\partial\theta}\left(\sin\theta\frac{\partial\psi}{\partial\theta}\right) = \frac{1}{r^2}\left(\frac{\cos\theta}{\sin\theta}\frac{\partial\psi}{\partial\theta} + \frac{\partial^2\psi}{\partial\theta^2}\right)\), a azymutalna ma w mianowniku \(r^2 \sin^2\theta\).
Daj znać, jak poszły testy lub jeśli gdzieś po drodze pojawi się rozbieżność!
Jeśli po przepisaniu kodu będziesz chciał:
• Zweryfikować pochodne analityczne dla wyższych orbitali (np. l=3)
• Przetestować zbieżność energii dla różnych stanów kwantowych
Wklej nowe wyniki lub fragment kodu, a pomogę sprawdzić.
Odpowiedzi wygenerowane przez AI mogą zawierać błędy. Więcej informacji
*/
