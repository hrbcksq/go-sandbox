package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var c complex64
	var p uintptr
	var r rune
	fmt.Printf("%v %v %v %q %v %v %v\n", i, f, b, s, c, p, r)
	// 0 0 false ""
}
