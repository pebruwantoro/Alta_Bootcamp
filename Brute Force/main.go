package main

import "fmt"

func main() {
	fmt.Println(MinMax([]int{5, 1, 2, 2, 14, 11, 2, -1}))
	fmt.Println(MinMax([]int{-100, -23, -9, -1, -1000, -7}))
}

func MinMax(input []int) (int, int) {
	// complete search == linear search, timecomplexity O(n)
	var min, max int
	if len(input) == 0 {
		return 0, 0
	}
	if len(input) > 0 {
		min = input[0]
		max = input[0]
	}
	for i := 1; i < len(input); i++ {
		if min > input[i] {
			min = input[i]
		}
		if max < input[i] {
			max = input[i]
		}
	}
	return min, max
}
