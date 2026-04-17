package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	longest := 0
	lastIndex := make([]int, 128)

	left := 0
	for right := range(len(s)) {
		currentChar := s[right]
		if lastIndex[currentChar] > left {
			left = lastIndex[currentChar]
		}
		longest = max(right-left+1, longest)
		lastIndex[currentChar] = right + 1
	}

	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
