package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // creates a slice with non-zero length

	// by default, a new slice's capacity == length, but we can pass a capacity
	// as an additional parameter to make
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("apd d:", s)
	s = append(s, "e", "f")
	fmt.Println("apd ef:", s)

	// slices can also be copied
	// we make a slice c with the same length as s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice[low:high] (e.g., s[2], s[3], s[4] here)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// up to but excluding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// up from and including s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and init a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices package
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// can be composed into multi-dimensional data structures
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	task1()
	task2()
	task3()
}

// Task 1: Slices share underlying memory.
// Modify index 0 of `sub` to 99, and predict what `original` will print.
func task1() {
	original := []int{10, 20, 30, 40, 50}
	fmt.Println("original:", original)
	sub := original[1:4] // [20, 30, 40]
	sub[0] = 99          // a view that mutates original
	fmt.Println("sub:", sub)
	original[0] = 99

	fmt.Println("original after:", original)
}

// Task 2: Fix the copy bug.
// `copy(dst, src)` only copies min(len(dst), len(src)) elements.
// Fix the declaration of `dst` using `make` so it successfully copies all elements from `src`.
func task2() {
	src := []int{1, 2, 3, 4}
	// var dst []int // Bug: len(dst) is 0, so copy(dst, src) copies 0 elements
	dst := make([]int, len(src))

	copy(dst, src)

	fmt.Println("dst:", dst)
}

// Task 3: Capacity & Reallocation (Mental Check).
// Predict whether a[0] changes when b[0] is mutated.
func task3() {
	a := make([]int, 2, 2) // len=2, cap=2 -> [0, 0]
	a[0] = 1
	a[1] = 2

	// `a` is at max capacity (2). Appending allocates a NEW backing array for `b`.
	b := append(a, 3)
	b[0] = 99

	// Predict output: will a[0] be 1 or 99?
	fmt.Println("a[0]:", a[0])
	fmt.Println("b[0]:", b[0])
}
