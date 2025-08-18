package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the map program")
	languageMap := make(map[string]string)
	languageMap["en"] = "English"
	languageMap["fr"] = "French"
	languageMap["es"] = "Spanish"

	fmt.Println("Language Map: ", languageMap)
	fmt.Println("Language for 'en': ", languageMap["en"])

	for key, value := range languageMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	delete(languageMap, "fr")
	fmt.Println("Language Map after deletion: ", languageMap)

}
