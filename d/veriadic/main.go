package main

import (
	"fmt"
)

func main() {
	//total := variadic(1, 2, 3, 4)
	//fmt.Println(total)
	//example2("a", 1, 2, 3, 4)
	//example3(1, 1, 2, 3, 4)
	//example4(1, "a", 6, 7, 8, 9)

	val := []int{6, 7, 8, 9}
	example4(1, "b", val...)
}

// ...(three dots)
// only single type allowed not allowedvariadic(val ...int, tp ...string)
func variadic(val ...int) int {
	total := 0
	for _, v := range val {
		total += v
	}
	return total
}

// veriadic parameter always be last parameter
func example2(a string, val ...int) {
	fmt.Println(a)
	fmt.Println(val)
}

func example3(a int, val ...int) {
	fmt.Println(a)
	fmt.Println(val)
}

func example4(a int, b string, val ...int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(val)
}
