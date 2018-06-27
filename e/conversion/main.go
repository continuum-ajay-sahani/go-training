package main

import (
	"fmt"
	"strconv"
)

func main() {
	//example1()
	//example2()
	//example3()
	example4()
}

func example1() {
	var x = 12
	var y = 12.1230123
	z := float64(x)
	fmt.Printf("Type of X=%T\n", z)
	fmt.Println(y + z)
}

func example2() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
	//	fmt.Println("I have this many: ", strconv.Itoa(x), x)
}

func example3() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
}

func example4() {

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)

}
