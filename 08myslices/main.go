package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcom to slices program")
	var fruitList = []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Println("Fruit List: ", fruitList)

	fruitList = append(fruitList, "Fig")
	fmt.Println("Updated Fruit List: ", fruitList)
	fmt.Println("Length of Fruit List: ", len(fruitList))

	fruitList = append(fruitList[:2], fruitList[3:]...) // Remove "Cherry"
	fmt.Println("After removing Cherry: ", fruitList)

	highScore := make([]int, 4)

	highScore[0] = 100
	highScore[1] = 200
	highScore[2] = 300
	highScore[3] = 400
	fmt.Println("High Scores: ", highScore)
	highScore = append(highScore, 500)
	fmt.Println("Updated High Scores: ", highScore)
	sort.Ints(highScore)
	fmt.Println("Sorted High Scores: ", highScore)
}
