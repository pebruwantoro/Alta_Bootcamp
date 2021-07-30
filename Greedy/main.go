package main

import "fmt"

func main() {
	fmt.Println(coinChange(42, []int{25, 10, 5, 1}))
	fmt.Println(coinChange(95, []int{25, 20, 10, 5}))
	fmt.Println(coinChange(53, []int{25, 15, 10, 5, 2, 1}))
	fmt.Println(coinChange(27, []int{15, 10, 5, 3, 2, 1}))
	fmt.Println(coinChange(675, []int{250, 200, 100, 50, 25}))
}

func coinChange(target int, coinsInput []int) []int {
	var result []int
	balanced := target

	for balanced != 0 {
		a := balanced
		c := -1
		for i := 0; i < len(coinsInput); i++ {
			coins := coinsInput[i]
			if balanced-coins >= 0 && balanced-coins < a {
				a = balanced - coins
				c = coins
			}
		}
		result = append(result, c)
		balanced = a
	}

	return result
}
