package main

import "fmt"

func isPalindrome(x int) bool {
	// According to the problem statement negative numbers cannot be palindrome
	if x < 0 {
		return false
	}
	digits := make([]int, 0)
	for x != 0 {
		digits = append(digits, x%10)
		x = x / 10
	}
	left, right := 0, len(digits)-1
	for left <= right {
		if digits[left] != digits[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
