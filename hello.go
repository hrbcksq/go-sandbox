package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	var result string
	if x < 0 {
		result = fmt.Sprint(complex(0, math.Sqrt(-x)))
	} else {
		result = fmt.Sprint(math.Sqrt(x))
	}
	return result
}

func main() {
	sqrt1 := sqrt(40)
	sqrt2 := sqrt(-2)
	fmt.Println(sqrt1, sqrt2)
}
