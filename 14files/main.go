package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := ("This is a placeholder for the main function in 14files/main.go")

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)
	defer file.Close()
}
