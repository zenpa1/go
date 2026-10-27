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

	// _ is the blank identifier
	// acts as write-only placeholder or "trash can" for values to be ignored
	// "I know a value is returned here, but I want to discard it" as Go
	// prohibits unused variables
	_, c := vals() // subset of returned values
	// In other words, it is positional, meaning you:
	// - Discard 1st value (3)
	// - Capture 2nd value (7)
	fmt.Println(c)

	q, ok := divide(10, 2)
	fmt.Println("10 / 2 =", q, "success:", ok)
	fmt.Println("Is 10 divisible by 0?", isDivisible(10, 0))
}

// Task 1: Complete 'divide' so it returns two values:
// 1. The result of integer division (a / b)
// 2. A boolean: true if division was successful (b != 0), or false if dividing by zero (b == 0).
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

// Task 2: Complete 'isDivisible' using the 'divide' function.
// Call 'divide(a, b)', but use the blank identifier `_` to discard the quotient
// and only return the boolean success status.
func isDivisible(a, b int) bool {
	_, result := divide(a, b)
	return result
}
