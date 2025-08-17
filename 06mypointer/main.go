package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer program")
	var ptr *int
	fmt.Println("Value of pointer is: ", ptr)

	myNumber := 100
	ptr = &myNumber
	fmt.Println("Value of pointer is: ", ptr)
}
