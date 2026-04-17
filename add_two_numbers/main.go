package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		head := &ListNode{
			Val:  0,
			Next: nil,
		}
		curr := head
		var quotiant, sum int
		for l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + quotiant
			quotiant = 0
			node := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
			l1 = l1.Next
			l2 = l2.Next

			if sum > 9 {
				quotiant = 1
			}
		}
		for l1 != nil {
			sum = l1.Val + quotiant
			quotiant = 0
			node := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
			l1 = l1.Next
			if sum > 9 {
				quotiant = 1
			}
		}
		for l2 != nil {
			sum = l2.Val + quotiant
			quotiant = 0
			node := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
			l2 = l2.Next
			if sum > 9 {
				quotiant = 1
			}
		}
		if sum > 9 {
			curr.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
		return head.Next
}

func arrToList(in []int) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	curr := head
	for _, v := range in {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		curr.Next = node
		curr = curr.Next
	}
	return head.Next
}

func listToArray(head *ListNode) []int {
	out := make([]int, 0)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func main() {
	fmt.Println(listToArray(addTwoNumbers(arrToList([]int{2, 4, 3}), arrToList([]int{5, 6, 4}))))
	fmt.Println(listToArray(addTwoNumbers(arrToList([]int{0}), arrToList([]int{0}))))
	fmt.Println(listToArray(addTwoNumbers(arrToList([]int{5}), arrToList([]int{5}))))
}
