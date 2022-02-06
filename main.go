package main

import "fmt"

func main() {
	// declaring a pointer with string type
	var firstNamePointer *string = new(string)
	// dereferencing operator
	*firstNamePointer = "Fred"
	fmt.Println(firstNamePointer, *firstNamePointer)

	firstName := "Fred1"
	// address of operator
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	// changing the value of the variable and testing the pointer and dereferencing of the pointer
	firstName = "Fred2"
	fmt.Println(ptr, *ptr)
}
