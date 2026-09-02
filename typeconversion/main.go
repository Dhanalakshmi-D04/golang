package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 10
	var y float64 = float64(x)
	fmt.Println(x)
	fmt.Println(y)
	main1()
}

func main1() {
	a := 10
	b := float64(a)
	fmt.Println(b)
	c := 10.75
	d := int(c)
	fmt.Println(d)
	e := 65
	f := string(rune(e))
	fmt.Println(f)
	g := 69
	h := strconv.Itoa(g)
	fmt.Println(h)
}
