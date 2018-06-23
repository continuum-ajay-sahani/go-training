package main

import (
	"fmt"
)

func main() {
	checkAddress()
	getInput()
}

func checkAddress() {
	sum := 150
	fmt.Println("sum:", sum)
	fmt.Println("address:", &sum)
}

func getInput() {
	a := 40
	var no int
	fmt.Println("Add no to be added")
	fmt.Scan(&no)
	total := a + no
	fmt.Println("Total=", total)
}
