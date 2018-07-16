package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What time is it?")
	now := time.Now().Hour()
	switch {
	case now < 8 && now > 21:
		fmt.Println("Good night")
	case now < 12:
		fmt.Println("Good morning")
	case now < 14:
		fmt.Println("Good afternoon")
	case now < 18:
		fmt.Println("Good evening")
	default:
		fmt.Println("Too far away.")
	}
}
