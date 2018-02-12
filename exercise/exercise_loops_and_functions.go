package exercise

import (
	"fmt"
	"math"
)

func Sqrt(x, eps float64) float64 {
	z := float64(1.0)
	tmp := float64(1.0)
	for tmp > eps {
		if z*z-x > 0 {
			tmp = z*z - x
		} else {
			tmp = x - z*z
		}
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func MainLoopAndFunction() {
	fmt.Printf("sqrt(2) = %.6f, eps=%.6f\n", Sqrt(2, 0.000001), math.Sqrt(2)-Sqrt(2, 0.000001))
}
