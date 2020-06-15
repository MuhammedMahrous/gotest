package main

import "fmt"
func main() {
	s := make([]int,4)
	fmt.Printf("s cap = %v and len = %v \n",cap(s),len(s))
	s = s[0:2]
        fmt.Printf("s cap = %v and len = %v \n",cap(s),len(s))
}
