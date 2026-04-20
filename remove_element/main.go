package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, 0
	for ; right < len(nums); right++ { 
		if nums[right] == val {
			continue
		}
		nums[left] = nums[right]
		left++
	}
	return left
}

func main() {
	nums := []int{3,3}
	fmt.Println(removeElement(nums, 5))
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
