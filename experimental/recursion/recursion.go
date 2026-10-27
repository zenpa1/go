package main

import "fmt"

// Declared at package level, so function body can reference fact directly
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Declare function signature variable FIRST
	var fib func(n int) int

	// Assign closure body to variable
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2) // Now fib is in scope
	}

	// fib := func(n int) int {
	// 	return fib(n-1) + fib(n-2) // ❌ Compiler Error: undefined: fib
	// }

	fmt.Println(fib(7))

	var sum func(n int) int
	sum = func(n int) int {
		if n <= 0 {
			return 0
		}

		return n + sum(n-1)
	}

	fmt.Println(sum(5))
}
