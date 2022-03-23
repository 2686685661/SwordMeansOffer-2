/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 12:35 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _21_删除链表的倒数第_n_个结点


/**
快慢指针
https://leetcode-cn.com/problems/c32eOV/solution/tu-jie-kuai-man-zhi-zhen-ji-qiao-yuan-li-rdih/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}

	pHead1, pHead2 := head, head

	h := &ListNode{}
	h.Next = head
	var pre = h

	for i := 0; i < n; i++ {
		pHead2 = pHead2.Next
	}
	for pHead2 != nil {
		pre = pHead1
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
	}

	pre.Next = pHead1.Next
	pHead1.Next = nil
	return h.Next

}