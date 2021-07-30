package main

import (
	"fmt"
)

func main() {
	fmt.Println(Exponential(3, 1))
	fmt.Println(Exponential(4, 0))
	fmt.Println(Exponential(7, 4))
	fmt.Println(Exponential(9, 5))
	fmt.Println(Exponential(10, 4))
}

func Exponential(x, y int) int {
	counter := y - 1
	result := x
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	}
	for counter > 0 {
		result *= x
		counter--
	}
	return result
}
