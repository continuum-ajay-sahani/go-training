package main

import "fmt"

func main() {
	//example1()
	example2()
}

func example1() {
	var name interface{} = "Pune"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func example2() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6)
}
