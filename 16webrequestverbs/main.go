package main


func main(){
	fmt.Println("Welcome to web request verbs")
}

func PerformPostJsonRequest(){
	const url ="http://localhost:8000"

	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go",
		"price":0,
		"platform":"learnCodeOnline.in"
	}
	`)

	response, err := http.Post(url,"application/json", requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}