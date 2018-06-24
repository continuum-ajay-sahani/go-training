package main

import (
	"fmt"
)

func main() {
	//example1()
	//example2()
	example3()
}

func example1() {
	var day = make(map[string]string)
	// day:=make(map[string]string)
	day["mon"] = "Day1"
	day["tue"] = "Day2"

	fmt.Println("Len=", len(day))

	for key, value := range day {
		fmt.Println(key, value)
	}
}

func example2() {
	day := map[string]string{
		"mon": "Day1",
		"tue": "Day2",
		"wed": "Day3",
	}
	fmt.Println(day)
	delete(day, "tue")
	fmt.Println(day)
}

func example3() {
	day := map[string]string{
		"mon": "Day1",
		"tue": "Day2",
		"wed": "Day3",
	}

	if val, ok := day["tue"]; ok {
		fmt.Println(val)
	}
}
