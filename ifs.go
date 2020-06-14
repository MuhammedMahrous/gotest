package main

import "fmt"

func main() {
	if true {
		fmt.Println("Indeed this is true")
	}

	if 9 > 1 {
		fmt.Println("Sure 9 is > 1 !")
	}

	// if with short statement
	if i := 0; i == 0 {
		fmt.Println("Well, you initialzed it with zero!")
		fmt.Printf("And it's still accessible here! %v \n", i)
	}

	// No you can't!
	//	fmt.Printf("Can I use i here? i= %v \n",i)

	// But you can use it inside 'else'
	if j := 0; j < 0 {
		fmt.Println("Can't leave this empty!")
	} else {
		fmt.Printf("j=%v is accessible inside 'else' \n",j)
	} 

}
