package main

import "fmt"

func lengthOfLastWord(s string) int {
	left, right := 0, len(s)
	pos := 0
	for s[right-1] == ' ' {
		right--
	}

	for ; left < right; left++ {
		if s[left] == ' ' {
			pos = left+1
		}
	}
	return right - pos
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("a"))
}
