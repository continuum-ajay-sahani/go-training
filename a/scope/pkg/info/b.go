package info

import (
	"fmt"
)

func samePkgCall() {
	fmt.Println("accessible in same package only")
	fmt.Printf("%s\n", inside)
}
