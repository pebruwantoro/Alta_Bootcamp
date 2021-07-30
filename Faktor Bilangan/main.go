package main

import "fmt"

func main() {
	fmt.Println(faktorBilangan(4))
	fmt.Println(faktorBilangan(7))
	fmt.Println(faktorBilangan(10))
	fmt.Println(faktorBilangan(13))
	fmt.Println(faktorBilangan(16))
	fmt.Println(faktorBilangan(20))
	fmt.Println(faktorBilangan(50))
}

func faktorBilangan(x int) []int {
	var result []int
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
