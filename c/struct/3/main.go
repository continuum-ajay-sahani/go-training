package main

import (
	"fmt"
)

func main() {
	//fieldOverriding()
	methodOverriding()
}

// Person info
type Person struct {
	First string
	Last  string
}

// Message msg
func (p Person) Message() {
	msg := "Person Message"
	fmt.Println(msg)
}

// PersonV2 info
type PersonV2 struct {
	Person
	First string
	Age   int
}

// Message msg
func (p2 PersonV2) Message() {
	fmt.Println("PersonV2 Message")
}

func fieldOverriding() {

	p1 := Person{
		First: "ajay",
		Last:  "sahani",
	}

	p2 := PersonV2{
		Person: p1,
		First:  "Kabir",
		Age:    29,
	}
	fmt.Println("First:", p2.First)
	fmt.Println("Last:", p2.Last)
	fmt.Println("Age:", p2.Age)
}

func methodOverriding() {
	p1 := Person{
		First: "ajay",
		Last:  "sahani",
	}

	p2 := PersonV2{
		Person: p1,
		First:  "Kabir",
		Age:    29,
	}

	p2.Message()
}
