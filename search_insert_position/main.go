package main

import "fmt"

func searchInsert(nums []int, target int) int {
	pos := 0
	for ; pos < len(nums); pos++ {
		if nums[pos] >= target {
			break
		}
	}
	return pos
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{}, 1))
}
