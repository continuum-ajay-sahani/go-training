package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
}
