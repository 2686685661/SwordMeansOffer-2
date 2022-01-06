/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 11:28 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _76_删除链表中重复的结点

type ListNode struct{
	Val int
	Next *ListNode
}

func deleteDuplication( pHead *ListNode ) *ListNode {

	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	Head := &ListNode{}
	Head.Next =pHead

	pre, last := Head, Head.Next

	for last != nil {
		if last.Next != nil && last.Val ==last.Next.Val {
			for last.Next != nil && last.Val == last.Next.Val {
				last = last.Next
			}
			pre.Next = last.Next
			last = last.Next
		} else {
			pre = pre.Next
			last = last.Next
		}
	}
	return Head.Next
}