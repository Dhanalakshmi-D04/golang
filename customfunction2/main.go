package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	x, y := swap("hello", "world")
	fmt.Println(x, y)
}
