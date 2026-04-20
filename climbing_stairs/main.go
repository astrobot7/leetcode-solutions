package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		curr, prev = prev+curr, curr
	}
	return curr
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}
