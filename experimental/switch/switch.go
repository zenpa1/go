package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i { // switch encloses cases
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// can switch a value that comes from chained functions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// can switch with cases having conditions
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Type switch (inspect and extract underlying concrete type at runtime)
	whatAmI := func(i any) {
		// any is Go's alias for interface{} (empty interface)
		// i any means it can accept ANY TYPE at runtime
		switch t := i.(type) {
		// i.(type) can only be used inside a switch statement which queries
		// the Go runtime to ask "what is the actual concrete type stored
		// inside i right now?"
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
