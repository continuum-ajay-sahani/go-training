package db

import (
	"fmt"
)

// Message variable contain message
var Message = "DB package message"

// GetMessage return message
func GetMessage() {
	fmt.Println("I am in DB package")
}
