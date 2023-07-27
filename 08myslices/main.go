package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type os fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Gauva")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 934
	highScores[2] = 434
	highScores[3] = 634
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
}
