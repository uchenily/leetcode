# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        v1 = []
        v2 = []
        while l1:
            v1.append(l1.val)
            l1 = l1.next
        while l2:
            v2.append(l2.val)
            l2 = l2.next
        v1.reverse()
        v2.reverse()
        s1 = 0
        s2 = 0
        for i in v1:
            s1 = s1 * 10 + i
        for j in v2:
            s2 = s2 * 10 + j
        sum = s1 + s2
        sl = list(str(sum))
        sl = list(map(int, sl))
        sl.reverse()
        head = ListNode(0)
        p = head

        for v in sl:
            tmp = ListNode(v)
            p.next = tmp
            p = p.next
        return head.next
