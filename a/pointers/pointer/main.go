package main

import (
	"fmt"
)

func main() {
	//pointer()
	pointer2()
}

func pointer2() {
	a := 45
	b := &a
	fmt.Println(*b)
	*b = 40
	fmt.Println(a)
}

func pointer() {
	a := 45
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	fmt.Println(b)

	//de referensing
	fmt.Println(*b)
}
