package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-1899898))
	fmt.Println(reverse(-54321))
	fmt.Println(reverse(543210))
	fmt.Println(reverse(987654321))
	fmt.Println(reverse(100000192))
	fmt.Println(reverse(101010109))
}

func reverse(x int) int {
	//convert int to string
	xStr := []byte(strconv.Itoa(x)) // convert int to string but using byte
	//fmt.Println(xStr) // will show byte number
	//fmt.Printf("%T, %s", xStr)
	// the purpose of convert int to str is getting each digit of int

	var a []byte // variable to save byte xStr[i]
	for i := len(xStr) - 1; i >= 0; i-- {
		a = append(a, xStr[i])
	}
	if x < 0 {
		a = append([]byte("-"), a[:len(a)-1]...) // condition for input is (-) || x < 0
	}
	result, _ := strconv.Atoi(string(a))             // convert string to int
	if result > 2147483647 || result < -2147483647 { // data type int can't more than that input
		return 0
	}
	return result
}
