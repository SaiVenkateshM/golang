package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("welcome to defer")
}
