package main

import (
	"fmt"
)

func main() {
	var i int
	j := i // j is an int

	// i := 42           // int
	f := 3.142        // float64
	g := 0.886 + 0.5i // complex128

	// fmt.Printf("%T %v", i, f, g)
	fmt.Printf("Value: %v Type: %T\n", j, j)
	fmt.Printf("Value: %v Type: %T\n", f, f)
	fmt.Printf("Value: %v Type: %T\n", g, g)
}
