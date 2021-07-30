package main

import (
	"fmt"
)

func main() {
	fmt.Println(concatOddString([]string{"b", "a", "ab", "c", "a"}))
	fmt.Println(concatOddString([]string{"lam", "a", "ab", "aa", "bb"}))
	fmt.Println(concatOddString([]string{"a", "b", "c", "d", "e"}))
	fmt.Println(concatOddString([]string{"s", "a", "y", "a", "ee"}))
	fmt.Println(concatOddString([]string{"kur", "v", "a", "aa", "ee"}))
}

func concatOddString(array []string) string {
	var result string
	var newArray []string

	for i := 0; i < len(array); i++ {
		if len(array[i])%2 == 0 {
		} else {
			newArray = append(newArray, array[i])
		}
	}
	for j := 0; j < len(newArray); j++ {
		result += newArray[j]
	}
	return result
}
