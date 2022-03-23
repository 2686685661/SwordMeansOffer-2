/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 9:02 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _24_反转链表

type ListNode struct {
	Val int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}