package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	ashutosh := User{"ashutosh", "ashu@go.dev", true, 16}
	fmt.Println(ashutosh)
	fmt.Printf("ashutosh details: %v\n", ashutosh)
	fmt.Printf("Name is %v\n", ashutosh.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
