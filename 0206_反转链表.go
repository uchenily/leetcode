package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	// 将head所指向的结点放到子链表的最后

	// ptr := newHead.Next
	// pre := newHead
	// for ptr != nil {
	// 	pre = ptr
	// 	ptr = ptr.Next
	// }
	// pre.Next = head
	// head.Next = nil

	head.Next.Next = head // 直接找到位置, 而不是遍历子链
	head.Next = nil

	return newHead
}
