/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/27 2:14 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _78_合并排序链表

import (
	"container/heap"
)

/**
给定一个链表数组，每个链表都已经按升序排列。
请将所有链表合并到一个升序链表中，返回合并后的链表。
 */

type ListNode struct {
	Val int
	Next *ListNode
}



/**   -- 最小堆 实现 -- **/

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Top() interface{} {
	if len(h) > 0 {
		return h[0]
	}
	return nil
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[:n-1]
	return x
}


//使用 k 个指针分别指向链表的头节点，每次都从这 k 个节点中选取值最小的节点确定为合并后的链表的第一个节点。
// 然后将指向最小值节点的指针向后移动一步，再比较这 k 个节点中选取值最小的节点确定为下一个节点。重复以上过程，所有链表就会被合并。
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) < 2 {
		return lists[0]
	}

	front := &ListNode{}
	p := front

	h := &Heap{}
	heap.Init(h)
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(h, l)
	}

	for h.Len() > 0 {
		min := h.Top().(*ListNode)
		heap.Pop(h)
		if min.Next != nil {
			heap.Push(h, min.Next)
		}
		p.Next = min
		p = p.Next
	}
	p.Next = nil
	return front.Next
}



/**   -- 用归并排序 实现 -- **/

func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) < 2 {
		return lists[0]
	}

	return mergeUp2Down(lists, 0, len(lists) - 1)

}

// 丛上往下归并
func mergeUp2Down(lists []*ListNode, low, high int) *ListNode {
	if low >= high {
		return lists[low]
	}

	mid := (low + high) >> 1
	head1, head2 := mergeUp2Down(lists, low, mid), mergeUp2Down(lists, mid + 1, high)
	return mergeList(head1, head2)
}


func mergeList(head1, head2 *ListNode) *ListNode {
	front := &ListNode{}
	p := front
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			p.Next = head1
			p = p.Next
			head1 = head1.Next
		} else {
			p.Next = head2
			p = p.Next
			head2 = head2.Next
		}
	}

	for head1 != nil {
		p.Next = head1
		p = p.Next
		head1 = head1.Next
	}

	for head2 != nil {
		p.Next = head2
		p = p.Next
		head2 = head2.Next
	}

	return front.Next
}