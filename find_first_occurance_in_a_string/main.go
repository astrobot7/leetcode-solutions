package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := range len(haystack) - len(needle) + 1 {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("abc", "c"))
}
