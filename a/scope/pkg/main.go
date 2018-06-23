package main

import (
	"fmt"

	"github.com/ContinuumLLC/go-training/a/scope/pkg/info"
)

func main() {
	info.CallMe()
	fmt.Println(info.OutsidePackage)
}
