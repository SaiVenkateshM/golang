package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welocomeMessage := "Welcome to the user input program"
	fmt.Println(welocomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input)
}
