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
}
