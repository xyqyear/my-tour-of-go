package main

import (
	"fmt"
)

func fabs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func mySqrt10Times(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func mySqrtWithPrecision(x, precision float64) float64 {
	z := 1.0
	for delta := z*z - x; fabs(z*z-x) > precision; delta = z*z - x {
		z -= delta / (2 * z)
	}
	return z
}

func mySqrt(x float64) float64 {
	return mySqrtWithPrecision(x, 0.00000001)
}

func main() {
	fmt.Println(mySqrt(2), mySqrtWithPrecision(2, 0.000000000000001))
}
