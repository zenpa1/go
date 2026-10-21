package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// type is inferred
	var d = true
	fmt.Println(d)

	// var declared without initializations are zero-valued (e.g., int = 0)
	var e int
	fmt.Println(e)

	// := is shorthand for declaring and initializing a variable (only inside functions)
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
