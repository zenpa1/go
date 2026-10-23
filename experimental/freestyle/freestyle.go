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

	divisibleBy2 := func(integer int) {
		switch { // Implicitly 'switch true'
		case integer%2 == 0:
			fmt.Println(integer, "is even")
		default:
			fmt.Println(integer, "is odd")
		}
	}

	isTrue := func(boolean bool) {
		switch boolean { // Inspect the value of boolean
		case true:
			fmt.Println("The statement", boolean, "is true")
		case false:
			fmt.Println("The statement", boolean, "is false")
		}
	}

	divisibleBy2(10)
	divisibleBy2(5)
	divisibleBy2(40)

	isTrue(true)
	isTrue(5-3 == 5)
	isTrue(2+2 == 4)

	checkTemp(31)
	checkTemp(30)
	checkTemp(15)
	checkTemp(14)

	checkStatus(200)
	checkStatus(404)
	checkStatus(500)
	checkStatus(1121)
}

// Task 1: Should print:
// "Hot" if temp > 30,
// "Moderate" if 15 <= temp <= 30
// "Cold" otherwise
func checkTemp(temp int) {
	switch {
	case temp > 30:
		fmt.Println(temp, "is Hot")
	case temp >= 15:
		fmt.Println(temp, "is Moderate")
	default:
		fmt.Println(temp, "is Cold")
	}
}

// Task 2: Should match HTTP status codes:
// 200 -> "OK"
// 404 -> "Not Found"
// 500 -> "Server Error"
func checkStatus(code int) {
	switch code {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Server Error")
	default:
		fmt.Println("Unknown Code")
	}
}

// Are you checking equality against fixed values (==)?
// ├── YES ──> Use `switch variable`
// └── NO  ──> (Ranges, <, >, &&, ||) ──> Use `switch` (Tagless)
