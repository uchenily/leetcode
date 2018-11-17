package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := head
	oddTail := head
	evenHead := head.Next
	evenTail := head.Next
	for oddTail != nil && evenTail != nil && oddTail.Next != nil && evenTail.Next != nil {
		oddTail.Next = oddTail.Next.Next
		oddTail = oddTail.Next

		evenTail.Next = evenTail.Next.Next
		evenTail = evenTail.Next
	}
	// 连接奇数链表和偶数链表
	if oddTail.Next == nil || evenTail.Next == nil {
		oddTail.Next = evenHead
	}
	return oddHead
}
