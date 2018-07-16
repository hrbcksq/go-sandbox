package main

import (
	"fmt"
)

func main() {
	sum := 0
	for sum < 1e3 {
		sum += 1
	}
	fmt.Println(sum)
}
