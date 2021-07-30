package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{4, 5, 9, 8, 1, 2, 6, 10}))
	fmt.Println(selectionSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
	fmt.Println(selectionSort([]int{10, 23, 100, 130, 200, 789, 1}))
	fmt.Println(selectionSort([]int{7, 9, 5, 6, 4, 3, 10, 11, 0, 1}))
	fmt.Println(selectionSort([]int{900, 8, 9, 3, 4, 1, 2}))
	fmt.Println(selectionSort([]int{0, 7, 6, 5, 1, 4}))
}

// this function is used to sort numbers from minimum to maximum
func selectionSort(arr []int) []int {
	// time complexity O(n^2)
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
