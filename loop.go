package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello %v \n", i)
	}

	i := 10
	for i > 0 {
		fmt.Printf("While like loop %v \n", i)
		i--
	}

	// make it true when trying it
	flag := false
	for flag {
		fmt.Println("LOOOOOOOOOOOOPING")
	}

	// another infite loop doing nothing
	/*
		for {

		}
	*/
}
