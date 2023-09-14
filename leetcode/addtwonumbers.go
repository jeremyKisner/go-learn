package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := l1
	first := []int{}
	for current != nil {
		fmt.Println(current.Val)
		first = append(first, current.Val)
		current = current.Next
	}
	reverseSlice(first)
	fmt.Println("first ", first)
	return &ListNode{}
}
