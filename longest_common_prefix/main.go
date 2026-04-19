package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	numStrs := len(strs)
	minLen := 1 << 10
	// Find the length of the shortest string
	for _, str := range strs {
		minLen = min(minLen, len(str))
	}
	var count int
	var mismatch bool
	for i := range minLen {
		for j := range (numStrs) - 1 {
			if strs[j][i] != strs[j+1][i] {
				mismatch = true
				break
			}
		}
		if !mismatch {
			count++
		}
	}
	return strs[0][:count]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
