/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 8:56 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _23_两个链表的第一个重合节点

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}