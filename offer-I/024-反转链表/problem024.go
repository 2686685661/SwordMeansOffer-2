/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 4:16 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _24_反转链表


type ListNode struct{
	Val int
	Next *ListNode
}

// 递归
func ReverseList( pHead *ListNode ) *ListNode {
	// write code here

	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	node := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil

	return node
}

// 循环
func ReverseList2( pHead *ListNode ) *ListNode {
	var h *ListNode = nil
	p := pHead

	for p != nil {
		t := p.Next
		p.Next = h
		h = p
		p = t
	}
	return h
}