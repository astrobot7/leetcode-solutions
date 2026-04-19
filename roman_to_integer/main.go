package main

import "fmt"

func romanToInt(s string) int {
	var val int
	var prev rune
	for i := range s {
		switch s[i] {
		case 'M':
			if prev == 'C' {
				val += 900
				continue
			}
			val += 1000
		case 'D':
			if prev == 'C' {
				val += 400
				continue
			}
			val += 500
		case 'C':
			if i != len(s)-1 && (s[i+1] == 'M' || s[i+1] == 'D') {
				prev = 'C'
				continue
			}
			if prev == 'X' {
				val += 90
				continue
			}
			val += 100
		case 'L':
			if prev == 'X' {
				val += 40
				continue
			}
			val += 50
		case 'X':
			if i != len(s)-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
				prev = 'X'
				continue
			}
			if prev == 'I' {
				val += 9
				continue
			}
			val += 10
		case 'V':
			if prev == 'I' {
				val += 4
				continue
			}
			val += 5
		case 'I':
			if i != len(s)-1 && (s[i+1] == 'X' || s[i+1] == 'V') {
				prev = 'I'
				continue
			}
			val += 1
		}
	}
	return val
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
