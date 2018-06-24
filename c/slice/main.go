package main

import (
	"fmt"
)

func main() {
	//example1()
	example2()
}

func example1() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	fmt.Println(months)

	summer := months[2:6]
	fmt.Println(summer)
}

func example2() {
	months := make([]string, 4)
	months[0] = "Jan"
	months[1] = "Feb"
	months[2] = "Mar"
	months[3] = "Apr"
	fmt.Println("Len=", len(months))
	fmt.Println("Cap=", cap(months))
	//months[4] = "May" //this will give error
	months = append(months, "May")
	fmt.Println("Len=", len(months))
	fmt.Println("Cap=", cap(months))
	fmt.Println(months)

	for index, value := range months {
		fmt.Println(index, value)
	}
}
