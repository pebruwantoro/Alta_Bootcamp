package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{2, 4, 6, 8, 10, 12, 14, 20}, 12))                               // true
	fmt.Println(binarySearch([]int{3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 27, 29, 53, 99}, 53))   // true
	fmt.Println(binarySearch([]int{}, 53))                                                         // false
	fmt.Println(binarySearch([]int{3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 27, 29, 53, 99}, 120))  // false
	fmt.Println(binarySearch([]int{3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 27, 29, 53, 99}, -123)) // false
}

func binarySearch(input []int, x int) bool {
	// time complexity of binary search is O(log n)
	if len(input) == 0 || x < input[0] || x > input[len(input)-1] {
		return false
	}
	if len(input) == 1 || x == input[0] {
		return true
	}
	pivot := len(input) / 2
	LeftInput := input[:pivot]
	RightInput := input[pivot:]
	fmt.Println(LeftInput)
	fmt.Println(RightInput)
	return binarySearch(LeftInput, x) || binarySearch(RightInput, x)
}
