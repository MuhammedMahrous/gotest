package main

import "fmt"

func main() {
	i := 55
	pi := &i
	fmt.Printf("%v is pointer value of i=%v \n", pi, i)
	/*
		 ./pointer.go:9:4: invalid operation: pi++ (non-numeric type *int)
		 pi++
		Yeah buddy, this is NOT C ... thank god
	*/
}
