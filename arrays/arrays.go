package main

import "fmt"

func main() {
	a := [2]string{"Hello", "World"}
	fmt.Println(a)
	var asliced []string = a[0:1]
	fmt.Println(asliced)
	asliced[0] = "Bye"
	fmt.Println("Now who's affected?")
	fmt.Println("Original array: ", a)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	// Still causes runtime error! out of bound
	// panic: runtime error: index out of range [6] with length 6
	fmt.Println(s[6])
}
