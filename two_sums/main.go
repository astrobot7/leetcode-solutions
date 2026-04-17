package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 99))
}

func twoSum(in []int, target int) []int {
	sumMap := make(map[int]int, len(in))
	for i, v := range in {
		diff := target - v
		if pos, ok := sumMap[v]; ok {
			return []int{pos, i}
		}
		sumMap[diff] = i
	}
	return nil
}
