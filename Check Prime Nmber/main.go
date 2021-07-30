package main

import "fmt"

func main() {
	fmt.Println(Prime(2))        // true
	fmt.Println(Prime(7))        // true
	fmt.Println(Prime(9))        // false
	fmt.Println(Prime(17))       // true
	fmt.Println(Prime(19))       // true
	fmt.Println(Prime(25))       // false
	fmt.Println(Prime(200))      // false
	fmt.Println(Prime(2017))     // true
	fmt.Println(Prime(20000021)) // false
}

func Prime(x int) bool {
	var a int
	if x <= 1 {
		return false
	}
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			a++
		}
	}
	if a == 2 {
		return true
	}
	return false
}
