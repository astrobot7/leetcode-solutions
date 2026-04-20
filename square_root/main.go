package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 2, x/2
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case mid == x/mid:
			return mid
		case mid < x/mid:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return right
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
