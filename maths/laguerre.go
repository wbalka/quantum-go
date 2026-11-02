package maths

import (
	"fmt"
)

type LaguerrePoly func(float64) float64

func GetLaguerrePolynomial(n int, m int) (LaguerrePoly, error) {
	switch n {
	case 1:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return -1.0
			}, nil
		}
	case 2:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return 2.0*ro - 4.0
			}, nil
		case 2:
			return func(ro float64) float64 {
				return 2.0
			}, nil
		}
	case 3:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return -3.0*ro*ro + 18.0*ro - 18.0
			}, nil
		case 2:
			return func(ro float64) float64 {
				return -6.0*ro + 18.0
			}, nil
		case 3:
			return func(ro float64) float64 {
				return -6.0
			}, nil
		}
	case 4:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return 4.0*ro*ro*ro - 48.0*ro*ro + 144.*ro - 96.0
			}, nil
		case 2:
			return func(ro float64) float64 {
				return 12.0*ro*ro - 96.0*ro + 144.0
			}, nil
		case 3:
			return func(ro float64) float64 {
				return 24.0*ro - 96.0
			}, nil
		case 4:
			return func(ro float64) float64 {
				return 24.0
			}, nil
		}
	case 5:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return -5.0*ro*ro*ro*ro + 100.0*ro*ro*ro - 600.0*ro*ro + 1200.0*ro - 600.0
			}, nil
		case 2:
			return func(ro float64) float64 {
				return -20.0*ro*ro*ro + 300.0*ro*ro - 1200.0*ro + 1200.0
			}, nil
		case 3:
			return func(ro float64) float64 {
				return -60.0*ro*ro + 600.0*ro - 1200.0
			}, nil
		case 4:
			return func(ro float64) float64 {
				return -120.0*ro + 600.0
			}, nil
		case 5:
			return func(ro float64) float64 {
				return -120.0
			}, nil
		}
	case 6:
		switch m {
		case 1:
			return func(ro float64) float64 {
				return 6.0*ro*ro*ro*ro*ro + 180.0*ro*ro*ro*ro - 1800.0*ro*ro*ro + 7200.0*ro*ro - 10800*ro - 4320.0
			}, nil
		case 2:
			return func(ro float64) float64 {
				return 30*ro*ro*ro*ro - 720.0*ro*ro*ro + 5400.0*ro*ro - 14400.0*ro + 10800.0
			}, nil
		case 3:
			return func(ro float64) float64 {
				return 120.0*ro*ro*ro - 2160.0*ro*ro + 10800.0*ro - 14400.0
			}, nil
		case 4:
			return func(ro float64) float64 {
				return 360.0*ro*ro - 4320.0*ro + 10800.0
			}, nil
		case 5:
			return func(ro float64) float64 {
				return 720.0*ro - 4320.0
			}, nil
		case 6:
			return func(ro float64) float64 {
				return 720
			}, nil
		}
	case 7:
		switch m {
		case 5:
			return func(ro float64) float64 {
				return -2520.0*ro*ro + 35280.0*ro - 105840.0
			}, nil
		case 6:
			return func(ro float64) float64 {
				return -5040.0*ro + 35280.0
			}, nil
		case 7:
			return func(ro float64) float64 {
				return -5040
			}, nil
		}
	}
	return nil, fmt.Errorf("Invalid combination of arguments n=%d, l=%d", n, m)
}
