/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 3:25 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _23_链表中环的入口结点

type ListNode struct{
	Val int
	Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	// 单一节点有环
	if pHead == pHead.Next {
		return pHead
	}

	// 确认是否有环
	po, pt := pHead, pHead.Next
	for pt != nil && pt.Next != nil {
		if po == pt {
			break
		}

		po = po.Next
		pt = pt.Next
		if pt != nil {
			pt =pt.Next
		}
	}
	if pt == nil || po != pt {
		return nil
	}

	// 存在环， 计算环的节点数
	count := 1
	pth := pt.Next
	for pth != pt {
		count += 1
		pth = pth.Next
	}

	// 确定环的入口
	po, pt = pHead, pHead
	for i := 0; i < count; i++ {
		pt = pt.Next
	}

	for po != pt {
		po = po.Next
		pt = pt.Next
	}

	return po

}