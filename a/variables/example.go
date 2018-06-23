package main

import (
	"fmt"
)

var a = "Hello"
var b, c int = 10, 5
var d, e = "Morning", true
var f float64

func main() {
	f = 1.234
	e := "Afternoon"
	var g = "test"

	fmt.Println("a-", a)
	fmt.Println("b-", b)
	fmt.Println("c-", c)
	fmt.Println("d-", d)
	fmt.Println("e-", e)
	fmt.Println("f-", f)
	fmt.Println("g-", g)

	// formtat specifier
	fmt.Printf("data type of a=%T\n", a)
	fmt.Printf("value of a = %v\n", a)
}
