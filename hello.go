package main

import (
	"fmt"
)

func main() {
	sum := 10
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
