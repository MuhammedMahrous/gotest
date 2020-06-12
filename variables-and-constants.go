package main

import "fmt"

// notice that you can declare multiple varialbes with one line
// vars withou inital value will take default type value; 0 for int and "" for string...etc
var a, b, c int

// And even initialize in one line
// initialized so type is omitted
var aInit, bInit, cInit = 1, 2, 3

// variables-and-constants.go:12:7: missing value in const declaration
//const Pi
const Pi = 3.14

func main() {
	var d string
	var dInit = "Initialized"
	fmt.Println("Not initialzed, huh!")
	fmt.Println(a, b, c, d)

	fmt.Println("Initialized")
	fmt.Println(aInit, bInit, cInit, dInit)

	// var inside a function? just use :=
	x, y, z := 4, 5, "hello!"
	fmt.Println(x, y, z)

	fmt.Println("Value of Pi is:", Pi)

	// ./variables-and-constants.go:30:5: cannot assign to Pi
	// Pi = 55
}
