package maths

import (
	"fmt"
	"math"
)

type SphericalHarm func(float64, float64) float64

/*
Spherical harmonics are the angular solutions when solving Laplace's equation in spherical coordinates
Substitution used: u=cosθ
*/
func GetSphericalHarm(l int, m int) (SphericalHarm, error) {
	switch l {
	case 0:
		switch m {
		case 0:
			return func(th float64, phi float64) float64 {
				return 1.0 / (2.0 * math.Sqrt(math.Pi))
			}, nil
		}
	case 1:
		switch m {
		case 0:
			return func(th float64, phi float64) float64 {
				return 0.5 * math.Sqrt(3/math.Pi) * math.Cos(th)
			}, nil
		case 1:
			return func(th float64, phi float64) float64 {
				return -0.5 * math.Sqrt(3/(2*math.Pi)) * math.Sin(th) * math.Cos(phi)
			}, nil
		case -1:
			return func(th float64, phi float64) float64 {
				return 0.5 * math.Sqrt(3/(2*math.Pi)) * math.Sin(th) * math.Sin(phi)
			}, nil
		}
	case 2:
		switch m {
		case 0:
			return func(th float64, phi float64) float64 {
				return 0.25 * math.Sqrt(5/math.Pi) * (3*math.Cos(th)*math.Cos(th) - 1)
			}, nil
		case 1:
			return func(th float64, phi float64) float64 {
				return 0.5 * math.Sqrt(15/math.Pi) * math.Cos(th) * math.Sin(th) * math.Cos(phi)
			}, nil
		case -1:
			return func(th float64, phi float64) float64 {
				return -0.5 * math.Sqrt(15/math.Pi) * math.Cos(th) * math.Sin(th) * math.Sin(phi)
			}, nil
		case 2:
			return func(th float64, phi float64) float64 {
				return 0.25 * math.Sqrt(15/math.Pi) * math.Sin(th) * math.Sin(th) * math.Cos(2*phi)
			}, nil
		case -2:
			return func(th float64, phi float64) float64 {
				return -0.25 * math.Sqrt(15/math.Pi) * math.Sin(th) * math.Sin(th) * math.Sin(2*phi)
			}, nil
		}
	case 3:
		switch m {
		case 0:
			return func(th float64, phi float64) float64 {
				return 0.25 * math.Sqrt(7/math.Pi) * (5*math.Cos(th)*math.Cos(th)*math.Cos(th) - 3*math.Cos(th))
			}, nil
		case 1:
			return func(th float64, phi float64) float64 {
				return 0.125 * math.Sqrt(42/math.Pi) * (5*math.Cos(th)*math.Cos(th) - 1) * math.Sin(th) * math.Cos(phi)
			}, nil
		case -1:
			return func(th float64, phi float64) float64 {
				return -0.125 * math.Sqrt(42/math.Pi) * (5*math.Cos(th)*math.Cos(th) - 1) * math.Sin(th) * math.Sin(phi)
			}, nil
		case 2:
			return func(th float64, phi float64) float64 {
				return 0.25 * math.Sqrt(105/math.Pi) * math.Sin(th) * math.Sin(th) * math.Cos(th) * math.Cos(2*phi)
			}, nil
		case -2:
			return func(th float64, phi float64) float64 {
				return -0.25 * math.Sqrt(105/math.Pi) * math.Sin(th) * math.Sin(th) * math.Cos(th) * math.Sin(2*phi)
			}, nil
		case 3:
			return func(th float64, phi float64) float64 {
				return 0.125 * math.Sqrt(70/math.Pi) * math.Sin(th) * math.Sin(th) * math.Sin(th) * math.Cos(3*phi)
			}, nil
		case -3:
			return func(th float64, phi float64) float64 {
				return -0.125 * math.Sqrt(70/math.Pi) * math.Sin(th) * math.Sin(th) * math.Sin(th) * math.Sin(3*phi)
			}, nil
		}
	}
	return nil, fmt.Errorf("Invalid combination of arguments n=%d, l=%d", l, m)
}
