package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	// Always make sure the bigger number is the first operand
	if len(b) > len(a) {
		a, b = b, a
	}
	out := ""
	operand := func(c byte) int {
		return int(c - '0')
	}
	var quotient int
	for i := len(b) - 1; i >= 0; i-- {
		op1 := operand(a[len(a)-len(b)+i])
		op2 := operand(b[i])
		sum := op1 + op2 + quotient
		if sum < 2 {
			out = strconv.Itoa(sum) + out
			quotient = 0
		} else {
			out = strconv.Itoa(sum%2) + out
			quotient = 1
		}
	}
	for i := len(a) - len(b) - 1; i >= 0; i-- {
		op1 := operand(a[i])
		sum := op1 + quotient
		if sum < 2 {
			out = strconv.Itoa(sum) + out
			quotient = 0
		} else {
			out = strconv.Itoa(sum%2) + out
			quotient = 1
		}
	}
	if quotient == 1 {
		out = "1" + out
	}
	return out
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("100", "110010"))
}
