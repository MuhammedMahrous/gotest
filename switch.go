package main

import "fmt"

func main() {
	makeNoise("Dog")
	makeNoise("Cat")
	makeNoise("Bird")
}

func makeNoise(animal string) {
	switch animal {
	case "Dog":
		fmt.Println("Ho Ho")
	case "Cat":
		fmt.Println("Mew Mew")
	default:
		fmt.Println("I don't know how to make such noise")
	}
}
