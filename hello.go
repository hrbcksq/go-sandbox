package main

import (
	"fmt"
)

func main() {
	var i, j int = 1, 2
	// the := short assignment statement can be used in place of a var declaration with implicit type
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, c, python, java)
}

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
