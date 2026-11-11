package main

import (
	"quantum-go/hydrogen"
)

/*
Below function runs benchmark calculations to find out the ideal set of parameters for different types of oprbitals
*/
func main() {
	/*parameters:
	main quantum number
	size of radial integration interval
	number of intervals
	lower limit of Gauss-Legendre quadrature points (program performs calulations for each set of G-L points to figure out the most accurate set)
	upper limit of Gauss-Legendre quadrature points
	*/
	hydrogen.CalculateSphericallySymmetricOrbitalEnergy(1, 2.0, 100, 14, 80)
	hydrogen.CalculateSphericallySymmetricOrbitalEnergy(2, 2.0, 100, 14, 80)
	hydrogen.CalculateSphericallySymmetricOrbitalEnergy(2, 1.0, 200, 14, 80)
	/*parameters:
	main quantum number
	azimuthal quantum number
	radial dimension of integration area
	number of areas
	lower limit of radial Gauss-Legendre quadrature points
	upper limit of radial Gauss-Legendre quadrature points
	radial G-L points number increment
	lower limit of angular G-L quadrature points
	upper limit of angular G-L quadrature points
	angular G-L points number increment
	*/
	hydrogen.CalculateAxiallySymmetricOrbitalEnergy(2, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	hydrogen.CalculateAxiallySymmetricOrbitalEnergy(3, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	hydrogen.CalculateAxiallySymmetricOrbitalEnergy(3, 2, 2.0, 100, 24, 40, 1, 24, 40, 1)
	hydrogen.CalculateAxiallySymmetricOrbitalEnergy(4, 1, 2.0, 100, 24, 40, 1, 24, 40, 1)
	hydrogen.CalculateAxiallySymmetricOrbitalEnergy(4, 2, 2.0, 100, 24, 40, 1, 24, 40, 1)
}
