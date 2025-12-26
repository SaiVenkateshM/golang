package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request verbs")
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostJsonRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go",
		"price":0,
		"platform":"learnCodeOnline.in"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Venkates")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
