package main

import (
	"fmt"
)

func main() {
	//example1()
	example2()
}

func example1() {
	defer func() {
		fmt.Println("inside defer")
	}()
	fmt.Println("tx1")
	fmt.Println("tx2")
}

func a() {
	fmt.Println("A")
}

func b() {
	fmt.Println("B")
}

func example2() {
	defer b()
	a()
}
