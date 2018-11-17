package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 这种解法对链表很长的情况会因为整数溢出不通过
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	n1 := 0
// 	n2 := 0
// 	for i := 0; l1 != nil; i++ {
// 		n1 += int(math.Pow10(i)) * l1.Val
// 		l1 = l1.Next
// 	}
// 	for i := 0; l2 != nil; i++ {
// 		n2 += int(math.Pow10(i)) * l2.Val
// 		l2 = l2.Next
// 	}
// 	sum := n1 + n2
// 	var head, tail *ListNode
// 	if sum == 0 {
// 		return &ListNode{Val: 0}
// 	}
// 	for sum > 0 {
// 		r := sum % 10
// 		ptr := &ListNode{Val: r}
// 		if head == nil {
// 			head = ptr
// 			tail = ptr
// 		} else {
// 			tail.Next = ptr
// 			tail = tail.Next
// 		}
// 		sum /= 10
// 	}
// 	// fmt.Println("n1", n1, "n2", n2)
// 	return head
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0 // 上一次是否有进位
	for l1 != nil && l2 != nil {
		p := &ListNode{Val: (l1.Val + l2.Val + carry) % 10}
		if head == nil {
			head = p
			tail = p
		} else {
			tail.Next = p
			tail = tail.Next
		}
		if l1.Val+l2.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// 加上剩下的部分
	for l1 != nil {
		p := &ListNode{Val: (l1.Val + carry) % 10}
		tail.Next = p
		tail = tail.Next
		if l1.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
	}
	for l2 != nil {
		p := &ListNode{Val: (l2.Val + carry) % 10}
		tail.Next = p
		tail = tail.Next
		if l2.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l2 = l2.Next
	}
	// 若最后有进位位, 添加上
	if carry == 1 {
		p := &ListNode{Val: 1}
		tail.Next = p
		tail = tail.Next
	}
	return head
}
