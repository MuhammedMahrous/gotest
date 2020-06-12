package main // since it's executable

import (
	"fmt"
	"time"
	"MuhammedMahrous/gotest/printingpkg"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello World This time")
	fmt.Println(time.Now())
	printingpkg.Newprint("Hi from pkg")
	fmt.Println(cmp.Diff("One word","the other"))
	fmt.Println(printingpkg.NewPrintable("Hello from NewPrintable"))
}
