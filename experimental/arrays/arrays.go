package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100 // set value
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])   // get value
	fmt.Println("len:", len(a)) // get length of array

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{6, 7, 8, 9, 10} // implicitly counts the number of elements for you
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // index 3 shall be 400
	fmt.Println("idx", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("2d: ", twoD)

	task1()
	task2()
}

// Task 1: Complete the declaration using the `[...]` syntax with index keying (`index: value`).
// It should create an array of size 6 where index 0 is 5, index 3 is 15, index 5 is 25,
// and all other indices remain zeroed. Then print its length using len().
func task1() {
	arr := [...]int{0: 5, 3: 15, 5: 25}

	fmt.Println("arr:", arr)
	fmt.Println("len:", len(arr))
}

// Task 2: Fix the compile error in this 2D matrix assignment.
// Pay attention to the matrix dimensions [rows][cols].
func task2() {
	var matrix [2][3]int

	matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("matrix:", matrix)
}
