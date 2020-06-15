package main

import "fmt"

func main() {
	// Define 1d array with size 5
	var array5 [5]int
	fmt.Println(array5)

	// Define 1d array with size not set; then later dynamically change the size with appending
	var array []int
	fmt.Printf("inital cap = %v and address of that array is %p \n",cap(array),&array)
	s := array[:]
	fmt.Printf("%p , %v, %v \n", s,cap(s),len(s))
	inc1array := append(s,1)
	fmt.Println(inc1array)
        fmt.Printf("after append cap = %v and address of that array is %p \n",cap(inc1array),inc1array)
        fmt.Printf("%p , %v, %v \n", s,cap(s),len(s))



        inc2array := append(inc1array,1,1)
        fmt.Println(inc2array)
        fmt.Printf("after append cap = %v and address of that array is %p \n",cap(inc2array),inc2array)



}
	
