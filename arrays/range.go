package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
		// This is a copy, you won't modify the original array
		v = 0
	}
	
	// Original array still intact
	fmt.Println(pow)
}
