package main

import (
	"fmt"
	"math"
)

func main() {
	a := "Hello"
	b := "World!"
	fmt.Println(a + " " + b)

	var c, d int = 67, 69
	fmt.Println(b, c, d)

	var e int
	fmt.Println("Hi! e is", e)

	// numeric constants have no type until given one via conversion
	// or through a context that requires one (e.g., variable assignment or function call)
	const forever_number = 420
	fmt.Println(forever_number)
	fmt.Println(float64(forever_number)) // explicitly converts
	fmt.Println(math.Sin(forever_number))
	// forever_number += 1 // isn't a valid target for assignment (+= is an ass. operator)

	const forever_true = true
	fmt.Println(forever_true)
	fmt.Println(!forever_true)

	for index := range 10 { // 0 to 10
		// fmt.Println("hello", index)
		if index%2 == 0 {
			fmt.Println(index, "even")
		} else {
			fmt.Println(index, "odd")
		}
	}
}
