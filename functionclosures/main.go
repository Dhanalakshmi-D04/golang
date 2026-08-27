package main

import "fmt"

func ASum() func() {
	var sum = 0
	return func() {
		sum += 1
		fmt.Println(sum)
	}
}

func main() {
	b := ASum()
	b()
	b()
	b()
	b()
	c := ASum()
	c()
}
