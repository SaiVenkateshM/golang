package main

import "fmt"

func main() {
	fmt.Println("Welcome to array program")
	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cherry"

	fmt.Println("Fruit List: ", fruitList)
}
