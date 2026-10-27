package main

import "fmt"

func vals() (int, int) { // returns 2 ints
	return 3, 7
}

func main() {
	// use the 2 different return values with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)

	_, c := vals() // subset of returned values
	fmt.Println(c)
}
