package main

import (
	"fmt"
)

func main() {
	//example1()
	example2()
}

func example1() {
	var a [5]int
	fmt.Println(a)
	fmt.Println(len(a))
	a[3] = 250
	fmt.Println(a)
}

func example2() {
	var a [20]int
	for i, v := range a {
		fmt.Println(i, v)
	}
}
