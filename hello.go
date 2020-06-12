package main // since it's executable

// factored style of imports
import (
	"MuhammedMahrous/gotest/printingpkg"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"time"
)

// but also this works and the mix too
import "math"

func main() {
	fmt.Println("Hello World This time")
	fmt.Println(time.Now())
	printingpkg.Newprint("Hi from pkg")
	fmt.Println(cmp.Diff("One word", "the other"))
	fmt.Println(printingpkg.NewPrintable("Hello from NewPrintable"))
	fmt.Println(math.Sqrt(14))
	result, msg := add(8, 9, "Adding 8 and 9")
	fmt.Println(result, msg)

	addNamed(1, 2, "Same as normal add()")
}

// not starting with capital letter so not exported
// same argument type can be compined in one declaration
// multiple results can be returned
func add(first, second int, msg string) (int, string) {
	fmt.Println(msg)
	return first + second, string("add() finished")
}

// named return
func addNamed(first, second int, msg string) (result int, resultmsg string) {
	fmt.Println(msg)
	result = first + second
	resultmsg = string("add() finished")
	return
}
