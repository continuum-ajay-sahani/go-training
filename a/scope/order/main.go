package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a")
}

func main() {

	//fmt.Println(x)
	// this will give error
	x := 5
	fmt.Println(x)
	a()
	test()
}

func test() {
	fmt.Println("Inside Test")
}
