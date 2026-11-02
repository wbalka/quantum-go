package maths

import "math"

type WaveFunction func(r float64, u float64) float64

func GetWaveFunction(n int, l int, m int, Z int) (WaveFunction, error) {
	lagPolly, _ := GetLaguerrePolynomial(n+l, 2*l+1)
	norm := NormalizationConstant(n, l, Z)
	sphericalHarm, _ := GetSphericalHarm(l, m)
	ro := 2.0 * float64(Z) / float64(n)
	return func(r float64, u float64) float64 {
		return norm * math.Pow(ro*r, float64(l)) * lagPolly(ro*r) * sphericalHarm(u) * math.Exp(-0.5*ro*r)
	}, nil
}
