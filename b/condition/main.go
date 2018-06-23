package main

import (
	"fmt"
)

func main() {
	a := "ajay"

	if a == "sanajay" {
		fmt.Println("sanjay")
	} else if a == "ajay" {
		fmt.Println("ajay")
	} else {
		fmt.Println("not found")
	}

	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	// food variable is not accessible here

	for i := 1; i < 20; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
