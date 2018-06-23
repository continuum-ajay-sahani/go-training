package main

import (
	"fmt"
)

func main() {
	//initCondition()
	//nested()
	//whileStyle()
	//noCondition()
	//breakLoop()
	continueLoop()
}

func initCondition() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}
}

func nested() {
	for i := 0; i < 4; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}

func whileStyle() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func noCondition() {
	for {
		fmt.Println("No Condition")
	}
}

func breakLoop() {
	i := 0
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}
}

func continueLoop() {

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 20 {
			break
		}
	}

}
