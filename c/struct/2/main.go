package main

import (
	"fmt"
)

func main() {
	example1()
	//example2()
}

// Person represent information of person
type Person struct {
	Name string
	Age  int
	Address
}

// BindMethod function binding
func (p Person) BindMethod() int {
	msg := p.addMethod()
	fmt.Println(msg)
	return 100
}

// Address represnt adddress
type Address struct {
	Area    string
	City    string
	State   string
	Pincode int
}

func (a Address) addMethod() string {
	return "address bind method"
}

func example1() {
	p := Person{
		Name: "Ajay",
		Age:  29,
		Address: Address{
			Area:    "Ghatkopar",
			City:    "Mumbai",
			State:   "Maharshtra",
			Pincode: 400084,
		},
	}

	fmt.Println(p.Name)
	fmt.Println(p.Area)
	fmt.Println(p.Pincode)
	fmt.Printf("%v\n", p)
	v := p.BindMethod()
	fmt.Println(v)
}

//---------------------------------------------

// Game holds common
type Game struct {
	Name string
}

// Cricket holds cricket
type Cricket struct {
	G           Game
	Description string
}

func example2() {
	gm := Game{
		Name: "Cricket",
	}
	c := Cricket{
		G:           gm,
		Description: "11 player",
	}
	fmt.Println(c.G.Name)
	fmt.Println(c.Description)
}
