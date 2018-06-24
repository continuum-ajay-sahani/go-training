package main

import "fmt"

func main() {
	example1()
	example2(2, 3)
	example3(2, 3)
	example4(2, "mon")
	example5(2, "tue")
	example6()
}

func example1() {
	fmt.Println("No args no return value")
}

func example2(a int, b int) int {
	sum := a + b
	fmt.Println("sum=", sum)
	return sum
}

func example3(a, b int) int {
	sum := a + b
	fmt.Println("sum=", sum)
	return sum
}

func example4(a int, b string) (int, string) {
	c := a + 5
	d := fmt.Sprintf("%s_test", b)
	return c, d
}

func example5(a int, b string) (c int, d string) {
	c = a + 5
	d = fmt.Sprintf("%s_test", b)
	return c, d
}

type point struct {
	x int
}

// method
func (p point) add(q point) int {
	return p.x + q.x
}

func example6() {
	p := point{
		x: 5,
	}

	q := point{
		x: 7,
	}

	fmt.Println(p.add(q))
}
