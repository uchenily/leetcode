#include<stdlib.h>

// Definition for singly-linked list.
struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    int lenA = 0;
    int lenB = 0;
    for (struct ListNode *pa = headA;pa!=NULL;pa=pa->next) {
        lenA++;
    }
    for (struct ListNode *pb = headB;pb!=NULL;pb=pb->next) {
        lenB++;
    }
    int diff = lenA-lenB; // 求出两个链表的长度差
    struct ListNode *pa=headA, *pb=headB;
    if (diff >= 0) { // A链表长
        for (;diff>0;diff--) {
            pa = pa->next;
        }
    } else { // B链表更长
        for (;diff<0;diff++) {
            pb = pb->next;
        }
    }
    for (;pa!=NULL&&pb!=NULL;pa=pa->next,pb=pb->next) {
        if (pa == pb) {
            return pa;
        }
    }
    return NULL;
}
