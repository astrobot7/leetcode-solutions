package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
}
