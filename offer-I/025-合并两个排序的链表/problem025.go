/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 8:01 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _25_合并两个排序的链表

type ListNode struct{
	Val int
	Next *ListNode
}

// 递归
func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}

	var node *ListNode = nil
	if pHead1.Val < pHead2.Val {
		node = pHead1
		node.Next = Merge(pHead1.Next, pHead2)
	} else {
		node = pHead2
		node.Next = Merge(pHead1, pHead2.Next)
	}
	return node
}

// 循环
func Merge2( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}

	var cur, head *ListNode = nil, nil
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			if head == nil {
				cur, head = pHead1, pHead1
			} else {
				cur.Next = pHead1
				cur = cur.Next
			}
			pHead1 = pHead1.Next

		} else {
			if head == nil {
				cur, head = pHead2, pHead2
			} else {
				cur.Next = pHead2
				cur = cur.Next
			}
			pHead2 = pHead2.Next
		}
	}
	if pHead1 == nil {
		cur.Next = pHead2
	} else {
		cur.Next = pHead1
	}
	return head
}