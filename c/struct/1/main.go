package main

import (
	"fmt"
)

// Person represent information of person
type Person struct {
	Name string
	Age  int
}

// GetInformation provide general information
func (p Person) GetInformation() string {
	return "Engineer"
}

func main() {

	p := Person{
		Name: "Ajay",
		Age:  29,
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.GetInformation())
}
