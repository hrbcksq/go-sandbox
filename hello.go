package main

import (
	"fmt"
)

// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.
const Pi = 3.14

func main() {
	const World = "Мир"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
