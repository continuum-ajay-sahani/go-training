package info

import (
	"fmt"
)

var inside = "Inside Package"

// OutsidePackage variable accessible from outside pkg
var OutsidePackage = "Outside Package"

// CallMe accessible from outside package
func CallMe() {
	fmt.Println("Accessible from outside package")
}
