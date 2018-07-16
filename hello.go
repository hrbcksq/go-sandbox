package main

import (
	"fmt"
	"math"
)

// Sqrt is the custom function
func Sqrt(x float64) float64 {
	epsilon := 1e-160
	z := 1.0
	for prev := 0.0; math.Abs(z-prev) > epsilon; {
		prev = z
		for index := 0; index < 10; index++ {
			z -= (z*z - x) / (2 * z)
		}
		fmt.Println(z)
	}

	return z
}

func main() {
	origin := 40.0
	result := Sqrt(origin)
	diff := math.Abs(math.Pow(result, 2) - origin)
	fmt.Printf("The result is: %v\nThe difference: is %v", result, diff)
}
