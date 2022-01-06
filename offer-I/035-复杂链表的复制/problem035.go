/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 2:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _35_复杂链表的复制

type RandomListNode struct {
	Label int
	Next *RandomListNode
	Random *RandomListNode
}

func Clone( head *RandomListNode ) *RandomListNode {
	if head == nil {
		return nil
	}

	// 复制每个节点，并将复制节点插在原节点之后
	cur := head
	for cur != nil {
		tmp := &RandomListNode{
			Label:  cur.Label,
		}
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}

	// 复制老节点的随机指针给新节点
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 拆分链表，将链表拆分为原链表和复制链表
	cur = head
	ch := head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
		}
		cur = cur.Next
	}
	return ch

}