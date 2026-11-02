package maths

import (
	"fmt"
	"math"
)

type SphericalHarm func(float64) float64

//We use substitution u=cosθ

func GetSphericalHarm(l int, m int) (SphericalHarm, error) {
	switch l {
	case 0:
		switch m {
		case 0:
			return func(u float64) float64 {
				return 1.0 / (2.0 * math.Sqrt(math.Pi))
			}, nil
		}
	case 1:
		switch m {
		case 0:
			return func(u float64) float64 {
				return 0.5 * math.Sqrt(3/math.Pi) * u
			}, nil
		}
	case 2:
		switch m {
		case 0:
			return func(u float64) float64 {
				return 0.25 * math.Sqrt(5/math.Pi) * (3*u*u - 1)
			}, nil
		}
	}
	return nil, fmt.Errorf("Invalid combination of arguments n=%d, l=%d", l, m)
}
