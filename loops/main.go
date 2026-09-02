package main

import (
	"fmt"
)

// for loop
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

func main1() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

// while loop

func main2() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

func main3() {
	numbers := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func main4() {
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
