package main

import "fmt"

func main() {
	fmt.Println("Welcome to the struct program")

	venkatesh := User{"Venkatesh", "saivenkatesh1403@gmail.com", true, 28}
	fmt.Println("Venkatesh details: ", venkatesh)
	fmt.Printf("Venkatesh details: %+v \n", venkatesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
