package main

import "fmt"

var Login string = "10000"

func main() {
	fmt.Println("variable")

	// string
	var stringVariable string
	fmt.Println(stringVariable)
	fmt.Printf("Variable is of type : %T \n", stringVariable)
	//checking the push to github.
	fmt.Printf("Login is of type : %T \n", Login)
}
