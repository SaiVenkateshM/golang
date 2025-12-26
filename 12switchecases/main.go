package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Switch and case program")

	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice value is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		fallthrough
	case 2:
		fmt.Println("Dice value is 2")
	default:
		fmt.Println("Defaulted value.")

	}
}
