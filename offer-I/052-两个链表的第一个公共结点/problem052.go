/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/15 12:54 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _52_两个链表的第一个公共结点


type ListNode struct{
   Val int
   Next *ListNode
}


func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	l1, l2 := getListLength(pHead1), getListLength(pHead2)
	if l1 == 0 || l2 == 0 {
		return nil
	}
	if l1 > l2 {
		pHead1 = walkStep(pHead1, l1 - l2)
	} else {
		pHead2 = walkStep(pHead2, l2 - l1)
	}

	for pHead1 != nil {
		if pHead1 == pHead2 {
			return pHead1
		}
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
	}
	return nil
}

func walkStep(pHead *ListNode, step int) *ListNode {
	for ; step > 0; step-- {
		pHead = pHead.Next
	}
	return pHead
}

func getListLength(pHead *ListNode) int {
	sum := 0
	if pHead == nil {
		return sum
	}

	for pHead != nil {
		sum += 1
		pHead = pHead.Next
	}
	return sum
}




/**
	假定 List1长度: a+n  List2 长度:b+n, 且 a<b
	那么 p1 会先到链表尾部, 这时p2 走到 a+n位置,将p1换成List2头部
	接着p2 再走b+n-(n+a) =b-a 步到链表尾部,这时p1也走到List2的b-a位置，还差a步就到可能的第一个公共节点。
	将p2 换成 List1头部，p2走a步也到可能的第一个公共节点。如果恰好p1==p2,那么p1就是第一个公共节点。  或者p1和p2一起走n步到达列表尾部，二者没有公共节点，退出循环。 同理a>=b.
	时间复杂度O(n+a+b)
 */
func FindFirstCommonNode2( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	p1, p2 := pHead1, pHead2
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = pHead2
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = pHead1
		}
	}
	return p1

}