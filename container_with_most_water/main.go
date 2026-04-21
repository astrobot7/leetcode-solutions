package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left <= right {
		area := (right-left) * min(height[left], height[right])
		maxArea = max(maxArea, area)
		if height[left] >= height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
