package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// requires explicit returns
// does not automatically return the value of the last expression
func concat(str1, str2 string) string {
	return str1 + " " + str2
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	str := concat("hello", "world")
	fmt.Println(str)

	res1 := multiplyThree(2, 2, 2)
	fmt.Println(res1)
	res2 := formatSum("S", 3, 7)
	fmt.Println(res2)
}

// Task 1: Refactor this function header using Go's parameter type deduplication syntax
// (collapse consecutive parameters of the same type), and fix the return logic so it compiles.
func multiplyThree(a, b, c int) int {
	return a * b * c
}

// Task 2: Fix the parameter types and return type signature.
// `prefix` should be a string, `x` and `y` should be ints, and it should return a string.
// (Hint: fmt.Sprintf("%s: %d", prefix, x+y) formats and returns a string)
func formatSum(prefix string, x, y int) string {
	return fmt.Sprintf("%s: %d", prefix, x+y)
}
