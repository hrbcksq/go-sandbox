package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Yesterday")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	case today + 4:
		fmt.Println("In four days")
	case today + 5:
		fmt.Println("In five days")
	default:
		fmt.Println("Too far away.")
	}
}
