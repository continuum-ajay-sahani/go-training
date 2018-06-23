package main

import (
	"errors"
	"fmt"
)

func main() {
	v, _ := getData()
	fmt.Println(v)

	v, err := getDataV2()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func getData() (int, error) {
	return 123, errors.New("Custom error")
}

func getDataV2() (int, error) {
	return 123, errors.New("Custom error")
}
