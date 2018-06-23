package main

import (
	"fmt"
)

var a int

func main() {
	name := "ajay"
	fmt.Println(name)
	{
		sum := 100
		fmt.Println(sum)
	}
	//sum is not accessible here
	increment()
	fmt.Println(a)
}

func increment() {
	a++
}
