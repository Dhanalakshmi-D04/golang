package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x, y := swap("hello", "world")
	fmt.Println(x, y)

	main1()
}

func swap1(a, b string) (string, string) {
	temp := a
	a = b
	b = temp

	return a, b
}

func main1() {
	a, b := swap1("he", "wo")
	fmt.Println(a, b)
}
