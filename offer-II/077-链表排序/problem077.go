/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/27 11:47 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _77_链表排序

import (
	"container/heap"
)

/**
要求：给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表
 */

type ListNode struct {
	Val int
	Next *ListNode
}





// 使用小顶堆进行排序

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

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := &Heap{}
	heap.Init(h)

	tmp := head

	for tmp != nil {
		heap.Push(h, tmp)
		tmp = tmp.Next
	}

	front := &ListNode{}
	cur := front

	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*ListNode)
		cur = cur.Next
	}
	cur.Next = nil

	return front.Next
}




// 归并排序：自上而下
// 时间复杂度O(nlogn)，空间复杂度O(logn)，即递归调用的深度
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head1, head2 := head, splitList(head)
	head1 = sortList2(head1)
	head2 = sortList2(head2)
	return mergeList(head1, head2)
}

// 获取链表中间节点，并从中间切断链表
func splitList(head *ListNode) *ListNode {
	slow, fast := head, head.Next

	for fast != nil  && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNode := slow.Next
	slow.Next = nil // 断尾
	return midNode
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

