package main

import "fmt"

func isValid(s string) bool {
	charQ := []rune{}

	for _, v := range s {
		switch v {
		case '(':
			charQ = append(charQ, ')')
		case '{':
			charQ = append(charQ, '}')
		case '[':
			charQ = append(charQ, ']')
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if len(charQ) > 0 && charQ[len(charQ)-1] == v {
				charQ = charQ[:len(charQ)-1]
			} else {
				return false
			}
		default:
		}
	}

	return len(charQ) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
}
