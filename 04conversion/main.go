package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welocomeMessage := "Welcome to the user input program"
	fmt.Println(welocomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input)

	numAge, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to number:", err)
		return
	} else {
		fmt.Println("Your age is: ", numAge-1)
		//fmt.Printf("Type of numAge is: %T\n", numAge)
	}

}
