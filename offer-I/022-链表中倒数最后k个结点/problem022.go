/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 2:57 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _22_链表中倒数最后k个结点


type ListNode struct{
	Val int
	Next *ListNode
}

func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
	pTemp := pHead
	for i := 0; i < k; i++ {
		if pTemp != nil {
			pTemp = pTemp.Next
		} else {
			return nil
		}
	}

	for pTemp != nil {
		pHead = pHead.Next
		pTemp = pTemp.Next
	}
	return pHead

}