package main

import "fmt"


func compute(executor func(arg1 string, arg2 string) string) string{
	return executor("Hello","World")
}

func main() {
	concat := func(arg1 string, arg2 string) string { return arg1 + " " + arg2 }

	fmt.Println(concat("Hi","There"))
	
	fmt.Println(compute(concat))
}
