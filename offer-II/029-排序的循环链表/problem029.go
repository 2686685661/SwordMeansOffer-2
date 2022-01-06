/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 12:14 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _29_排序的循环链表

type Node struct {
	Val int
	Next *Node
}

/**
指针p来遍历链表
1。 若输入的head为null,则直接新建对象,并将next指向自己。
2。因为链表是升序的,所以当p.val <= x <= p.next.val时,可以直接插入值而不破坏顺序
3。p.val > p.next.val，说明此时p指向了链表的最大值,而p.next则指向链表的最小值。
	若x >= p.val,即大于链表的最大值,则直接插入到p与p.next之间
	若x <= p.next.val,即小于链表的最小值,也直接插入到p与p.next之间
4。遍历一圈链表后,仍没有找到可插入的地方,则说明该链表所有值都是相同的,则x可以插入到任意地方
 */
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{
			Val:  x,
			Next: nil,
		}
		node.Next = node
		return node
	}

	p := aNode
	var t *Node
	h := func() {
		node := &Node{
			Val:  x,
			Next: p.Next,
		}
		p.Next = node
	}

	for p != t {
		if t == nil {
			t = aNode
		}
		if  (x >= p.Val && x <= p.Next.Val) || (p.Val > p.Next.Val && (x >= p.Val || x <= p.Next.Val))   {
			h()
			return aNode
		}
		p = p.Next
	}

	h()
	return aNode

}