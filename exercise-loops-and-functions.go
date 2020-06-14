package main

import (
	"fmt"
	"math"
)

const Precision = 0.000001

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs((z*z)-x) > Precision {
		z -= (z*z - x) / (2*z)
	}

	return z	
}

func main() {
	fmt.Println(Sqrt(2))
}
