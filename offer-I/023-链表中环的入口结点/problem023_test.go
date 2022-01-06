/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 3:41 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _23_链表中环的入口结点

import (
	"fmt"
	"testing"
)

func TestEntryNodeOfLoop(t *testing.T) {

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
	l5.Next = l3
	fmt.Printf("%+v\n", EntryNodeOfLoop(l1))
}