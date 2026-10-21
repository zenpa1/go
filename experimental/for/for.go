package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("=====")

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("=====")

	// Do this n times (range int)
	// range produces values to iterate over and := declares a new variable
	// to receive each value
	/*
		In this case:
			Iteration 1 i=0 -> range 0
			Iteration 2 i=1 -> range 1
			Iteration 3 i=2 -> range 2
	*/
	for index := range 3 {
		fmt.Println("range", index)
	}

	fmt.Println("=====")

	// `for` without a condition will loop repeatedly until you `break` out of the
	// loop or `return` from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("=====")

	for number := range 6 {
		// if n%2 == 0 {
		// 	fmt.Println(number, "even")
		// } else {
		// 	fmt.Println(number, "odd")
		// }

		if number%2 == 0 {
			continue // continues/skips to next iteration of loop
		}

		fmt.Println(number)
	}
}
