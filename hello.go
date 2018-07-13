package main

import (
	"fmt"
)

func swap(x, y int) (xx int, yy int) {
	xx = y
	yy = x
	return
}

func main() {
	a, b := swap(0, 1)
	fmt.Println(a, b)
}
