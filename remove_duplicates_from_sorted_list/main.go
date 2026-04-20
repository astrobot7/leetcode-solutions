package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	curr := head.Next
	for curr != nil {
		if curr.Val == prev.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head
}

func main() {
	fmt.Println(listToArray(deleteDuplicates(arrToList([]int{1, 1, 2}))))
	fmt.Println(listToArray(deleteDuplicates(arrToList([]int{1, 1, 2, 3, 3}))))
	fmt.Println(listToArray(deleteDuplicates(arrToList([]int{}))))
}
