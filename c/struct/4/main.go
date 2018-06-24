package main

import (
	"fmt"
)

func main() {
	p := &person{"Ajay", "Sahani", 29}
	fmt.Println(p.GetName())
	p.UpdateFirst()
	fmt.Println(p.GetName())
	p.UpdateLast()
	fmt.Println(p.GetName())
}

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) GetName() string {
	return fmt.Sprintf("%s %s", p.First, p.Last)
}

func (p person) UpdateFirst() {
	p.First = "Raj"
}

func (p *person) UpdateLast() {
	p.Last = "Kumar"
}
