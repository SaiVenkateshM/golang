package main

import "fmt"

var Login string = "10000"

func main() {
	fmt.Println("variable")

	// string
	var stringVariable string
	var isLoggedIn bool = true
	fmt.Println(stringVariable)
	fmt.Printf("Variable is of type : %T \n", stringVariable)
	fmt.Printf("Login is of type : %T \n", Login)
	fmt.Printf("isLoggedIn is of type : %T \n", isLoggedIn)
}
