package leetcode

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	ptr := head
	ptr1 := l1
	ptr2 := l2
	for ptr1 != nil || ptr2 != nil {
		var value int
		if (ptr1 != nil && ptr2 != nil && ptr1.Val < ptr2.Val) || ptr2 == nil {
			value = ptr1.Val
			ptr1 = ptr1.Next
		} else {
			value = ptr2.Val
			ptr2 = ptr2.Next
		}
		ptr.Next = &ListNode{Val: value, Next: nil}
		ptr = ptr.Next
	}
	return head.Next // 忽略头结点
}
