package maths

import "math"

type WaveFunction func(r float64, u float64) float64

/*
Function that returns the wave function for hydrogen-like atom orbital.
The wave function is produced from components: normalization constant, laguerre polynomial, spherical harmonic
ro is an auxilary variable
Parameters:
n - main quantum number
l - azimuthal quatnum number
m - magnetic quantum number
Z - number of protons in nucleus (total charge of nucleus)
*/
func GetWaveFunction(n int, l int, m int, Z int) (WaveFunction, error) {
	lagPolly, _ := GetLaguerrePolynomial(n+l, 2*l+1)
	norm := GetNormalizationConstant(n, l, Z)
	sphericalHarm, _ := GetSphericalHarm(l, m)
	ro := 2.0 * float64(Z) / float64(n)
	return func(r float64, u float64) float64 {
		return norm * math.Pow(ro*r, float64(l)) * lagPolly(ro*r) * sphericalHarm(u) * math.Exp(-0.5*ro*r)
	}, nil
}
