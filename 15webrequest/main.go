package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	content := string(databyte)
	fmt.Println("Content length is: ", len(content))
	fmt.Println("Content is: ", content)

}
