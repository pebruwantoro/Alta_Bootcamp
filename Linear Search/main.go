package main

import "fmt"

func main() {
	fmt.Println(linearSearch(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(linearSearch(9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(linearSearch(11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(linearSearch(7, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(linearSearch(25, []int{0, 5, 10, 15, 20, 25, 35, 40, 45, 50}))
}

// this function is used to check variable target is that in the array
func linearSearch(target int, arr []int) int {

	// time complexity O(n)
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			return i
		}
	}
	return 0
}
