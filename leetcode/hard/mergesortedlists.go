package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	numOfLists := len(lists)
	if numOfLists == 0 {
		return nil
	}
	curList := lists[0]
	if numOfLists == 1 {
		return curList
	}
	for i := 1; i < numOfLists; i++ {
		curList = mergeList(curList, lists[i])
	}
	return curList
}

func mergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
			current = current.Next
		} else {
			current.Next = l2
			l2 = l2.Next
			current = current.Next
		}
	}
	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}
	return head.Next
}
