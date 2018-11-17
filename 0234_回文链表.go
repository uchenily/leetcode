package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	fastPtr := head
	slowPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}
	// slowPtr指向了最中间的结点
	// 反转后半截链表
	ptr1 := reverseList(slowPtr)
	ptr2 := head
	// 现在的链表是这样的 []->[]->[]...[]...[]<-[]<-[]
	for ptr1 != nil {
		if ptr1.Val == ptr2.Val {
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		} else {
			return false
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
