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
}
