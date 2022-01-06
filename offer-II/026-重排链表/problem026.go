/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/29 5:59 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _26_重排链表

type ListNode struct {
	Val int
	Next *ListNode
}

// 通过slice，重建链表
func reorderList(head *ListNode)  {

	if head == nil || head.Next == nil {
		return
	}

	ns := make([]*ListNode, 0)

	for node := head; node != nil; node = node.Next {
		ns = append(ns, node)
	}

	i, j := 0, len(ns) - 1

	for i < j {
		ns[i].Next = ns[j]
		i++
		if i == j {
			break
		}
		ns[j].Next = ns[i]
		j--
	}
	ns[i].Next = nil

}



// 指定中间节点，将中间节点右边的子链表反转，合并链表
func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := getMiddleNode(head)
	l1, l2 := head, mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func getMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func mergeList(l1, l2 *ListNode) {
	l1t, l2t := l1, l2
	for l1 != nil && l2 != nil {
		l1t = l1.Next
		l2t = l2.Next

		l1.Next = l2
		l1 = l1t
		l2.Next = l1
		l2 = l2t
	}

}

