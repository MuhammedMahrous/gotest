package main

import "fmt"

func main() {
	defer fmt.Println("1.This will only be called whem main() is finished -1.")
	defer fmt.Println("2.You know what is weird? I was deferred last but I will execute first after main() \n I guess we're in a stack! -2.")
	fmt.Println("3.main() is saying Bye Bye, I wonder if something will happen afterwards! -3.")
}
