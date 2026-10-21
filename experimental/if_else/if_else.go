package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	} // Once execution exits the entire if-else block, num goes out of scope
	// which prevents variable name pollution and using a stale state

	// This syntax is most commonly used to capture a function call's return
	// value and inspect it (e.g., error handling)
	// if err := doSomething(); err != nil {
	// log.Println("Failed:", err)
	// return
	// }
	// 'err' doesn't leak into the rest of the function!

	// fmt.Println(num) // num exists only within the context of the if-else block
}
