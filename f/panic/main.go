package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Panic occured")
	}()
	example()
}

func example() {
	fmt.Println("Test")
	panic("err")
}
