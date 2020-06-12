package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	/*
		No implicit conversion even for values that might work; int to float
		./type-conversions.go:12:6: cannot use f (type float64) as type uint in assignment
		./type-conversions.go:13:6: cannot use x (type int) as type float64 in assignment

	*/
	//	var a uint = f
	//	var b float64 = x
	fmt.Println(x, y, z)
}
