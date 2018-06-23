package main

import (
	"fmt"
)

func sum() int {
	return 100
}

func main() {
	sum := sum()
	// Bad practice now sum is not function now it is variable
	fmt.Println(sum)
}
