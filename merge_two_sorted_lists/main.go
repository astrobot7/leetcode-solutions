package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{Val: -1, Next: nil}
	head := merged
	for list1 != nil && list2 != nil {
		node := &ListNode{
			Next: nil,
		}
		if list1.Val < list2.Val {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}
		merged.Next = node
		merged = merged.Next
	}
	for list1 != nil {
		merged.Next = &ListNode{
			Val: list1.Val,
			Next: nil,
		}
		merged = merged.Next
		list1 = list1.Next
	}
	for list2 != nil {
		merged.Next = &ListNode{
			Val: list2.Val,
			Next: nil,
		}
		merged = merged.Next
		list2 = list2.Next
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
	fmt.Println(listToArray(mergeTwoLists(arrToList([]int{1, 2, 4}), arrToList([]int{1, 3, 4}))))
	fmt.Println(listToArray(mergeTwoLists(arrToList([]int{}), arrToList([]int{}))))
	fmt.Println(listToArray(mergeTwoLists(arrToList([]int{}), arrToList([]int{0}))))
}
