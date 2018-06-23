package main

import (
	"fmt"
)

const template = "Hello"

const (
	a = 100
	b = true
	c = "Good Morning"
)

const (
	d = iota //0
	e = iota //1
	f = iota //2
)

const (
	g = iota
	h
	i
)

const (
	j = iota
	k = 100
	l = iota
)

func main() {
	var a = 5
	fmt.Println(template)
	a = 10
	fmt.Println(a)
	//template = "aja" //get error
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)
	fmt.Println(j, k, l)
}
