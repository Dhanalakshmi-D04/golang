package main

import (
	"fmt"
)

func main() {
	name := "Laksh"
	age := 20
	height := 5.4
	isStudent := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", height)
	fmt.Printf("%T\n", isStudent)
}
