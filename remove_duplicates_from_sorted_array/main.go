package main

import "fmt"

// removeDuplicates is a solution using maps and creating a new array, it's not efficient 
func removeDuplicates(nums []int) int {
	intMap := make(map[int]struct{})
	out := make([]int, 0, len(nums))
	count := 0
	for _, n := range nums {
		if _, ok := intMap[n]; ok {
			continue
		}
		out = append(out, n)
		intMap[n] = struct{}{}
		count++
	}
	copy(nums, out)
	return count
}

// removeDuplicatesWithSlidingWindow is the efficient solution but the final array may contain duplicate elements after the unique elements
func removeDuplicatesWithSlindingWindow(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		} else {
			left++
			nums[left] = nums[right]
		}
	}
	return left+1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicatesWithSlindingWindow([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
