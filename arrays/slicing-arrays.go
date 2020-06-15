package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println(s)
	s = s[:2]
	fmt.Printf("sliced to index 2 and capacity = %v and length = %v and slice = %v \n",cap(s),len(s),s)
	s = s[:]
	fmt.Println(s)
	
	var nils []int
	fmt.Printf(" nilS = %v and its capacity = %v and length = %v \n",nils,cap(nils),len(nils))
	if nils == nil {
		fmt.Println("Indeed Nil")
	}
}
