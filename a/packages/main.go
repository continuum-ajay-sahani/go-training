package main

import (
	"fmt"

	"github.com/ContinuumLLC/go-training/a/packages/db"
	"github.com/ContinuumLLC/go-training/a/packages/ui"
)

func main() {
	fmt.Println("I am main package.")
	db.GetMessage()
	ui.GetMessage()
}
