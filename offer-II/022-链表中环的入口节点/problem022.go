/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 1:56 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _22_链表中环的入口节点

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode-cn.com/problems/c32eOV/solution/jian-zhi-offer-ii-022-lian-biao-zhong-hu-8f1m/
func detectCycle(head *ListNode) *ListNode {

	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow =slow.Next
		if fast == slow {
			break
		}
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast

}