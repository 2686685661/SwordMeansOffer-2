/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/29 8:15 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _27_回文链表


type ListNode struct {
	Val int
	Next *ListNode
}

// 快慢指针法
// 寻找中心节点，反转中心右面链表，对比两链表相同
// 空间复杂度O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := middleNode(head)
	l1, l2 := head, mid.Next
	mid.Next = nil
	l2 = reverseListIteration(l2)
	return diffList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}


// 递归版本的反转链表，空间复杂度为O(n)，可以优化成循环实现
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 循环实现反转链表，空间复杂度为O(1)
func reverseListIteration(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func diffList(l1, l2 *ListNode) bool {

	if l1 == nil || l2 == nil {
		return false
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}


// 把值放到数组中，然后双指针判断
// 但是空间复杂度O(n)
func isPalindrome2(head *ListNode) bool {
	vals := make([]int, 0)
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	n := len(vals)

	for i:= 0; i< n / 2; i++ {
		if vals[i] != vals[n - 1 - i] {
			return false
		}
	}
	return true
}


