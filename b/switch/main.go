package main

import (
	"fmt"
)

func main() {
	//switchStatementNoMatch()
	//switchStatementMatch()
	//fallThrough()
	//multipleValue()
	noExpression()
}

func switchStatementNoMatch() {
	switch "Ajay" {
	case "Raj":
		fmt.Println("Raj Here")
	case "Vijay":
		fmt.Println("Vijay")
	default:
		fmt.Println("Nothing Match")

	}
}

func switchStatementMatch() {
	switch "Ajay" {
	case "Ajay":
		fmt.Println("Ajay Here")
	case "Vijay":
		fmt.Println("Vijay")
	default:
		fmt.Println("Nothing Match")

	}
}

func fallThrough() {
	switch "Ajay" {
	case "Ajay":
		fmt.Println("Ajay Here")
		fallthrough
	case "Ravi":
		fmt.Println("Ravi Here")
	case "Vijay":
		fmt.Println("Vijay")
	default:
		fmt.Println("Nothing Match")

	}
}

func multipleValue() {
	switch "Sanjay" {
	case "Ajay", "Sanjay", "Ovais":
		fmt.Println("Ajay,Sanjay,Ovais Here")
	case "Ravi", "Cricket":
		fmt.Println("Ravi Here")
	case "Vijay":
		fmt.Println("Vijay")
	default:
		fmt.Println("Nothing Match")

	}
}

func noExpression() {
	name := "Ajay"
	switch {
	case len(name) < 3:
		fmt.Println("Ravi Here")
	case name == "Ajay":
		fmt.Println("Ajay Here")
	case name == "Ravi":
		fmt.Println("Ravi Here")
	case name == "Vijay":
		fmt.Println("Vijay")
	default:
		fmt.Println("Nothing Match")

	}
}
