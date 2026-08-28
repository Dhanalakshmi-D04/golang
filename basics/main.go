package main

import "fmt"

func main() {
	a := [3]string{"a", "b", "c"}
	b := a[0:1]
	fmt.Println(b)

	b[0] = "laksh"
	fmt.Println(a)
}
