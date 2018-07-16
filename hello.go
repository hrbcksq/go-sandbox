package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// A deferred function's arguments are evaluated when the defer statement is evaluated.
func def() {
	i := 0
	defer fmt.Println("Defer function argumets are evalueated when the defer statement is evaluated: %v", i)
	i++
	return
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func counting() {
	defer fmt.Println()
	for index := 0; index < 5; index++ {
		defer fmt.Print(index)
	}
	defer fmt.Print("The counting function result is: ")
}

// Deferred functions may read and assign to the returning function's named return values.
func mutateNamedReturn() (i int) {
	defer func() { i++ }() // Anonymous function example
	return 1
}

func main() {
	counting() // defer calls 43210
	def()
	fmt.Printf("Defer named return values changes: %v", mutateNamedReturn())
}
