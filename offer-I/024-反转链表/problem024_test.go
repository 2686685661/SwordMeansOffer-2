/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 4:20 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _24_反转链表

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	n := ReverseList(l1)
	for n != nil {
		fmt.Printf("%+v\n", n)
		n = n.Next
	}
}
