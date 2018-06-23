package main

import (
	"fmt"
)

var a = "Hello"
var b, c int = 10, 5
var d, e = "Morning", true
var f float64

type myData float64

func main() {
	f = 1.234
	e := "Afternoon"
	var g = "test"

	var per myData
	per = 8.96

	fmt.Println("a-", a)
	fmt.Println("b-", b)
	fmt.Println("c-", c)
	fmt.Println("d-", d)
	fmt.Println("e-", e)
	fmt.Println("f-", f)
	fmt.Println("g-", g)
	fmt.Println(increment(5))
	fmt.Println(per)

	// formtat specifier
	fmt.Printf("data type of a=%T\n", a)
	fmt.Printf("value of a = %v\n", a)
	fmt.Printf("per type = %T\n", per)
}

var increment = func(a int) int {
	a++
	return a
}
