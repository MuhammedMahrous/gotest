package main

import "fmt"

type Vertex struct {
	x int
	y int
}
// Struct Literals to init a struct
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)


func main() {
	point := Vertex{15,9}
	fmt.Printf("So this is type '%T' with value = '%v' \n",point,point)
	fmt.Printf("And you can access x = %v \n",point.x)

	structPointer := &point
	fmt.Printf("Accessing x=%v with structPointer directly without using (*structPointer).x \n",structPointer.x)
	fmt.Println("Multiple ways to init a struct! => ", v1, p, v2, v3)

}
