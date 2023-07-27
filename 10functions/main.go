package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proResults, mymessage := proAdder(2, 3, 4, 5, 40)
	fmt.Println("Pro result is :", proResults, mymessage)

}
func proAdder(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}
	return total, "Hi Pro result "
}

func adder(n int, m int) int {
	return n + m
}

func greeter() {
	fmt.Println("Golang Interview")
}
