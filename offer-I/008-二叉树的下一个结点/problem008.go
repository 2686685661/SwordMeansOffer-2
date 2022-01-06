/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 3:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _08_二叉树的下一个结点

type TreeLinkNode struct {
	Val int
	Left *TreeLinkNode
	Right *TreeLinkNode
	Next *TreeLinkNode
}


/**
	如果一个节点有右子树，那么中序遍历的下一个节点就是它的右子树中的最左子节点
	如果一个节点没有右子树，且该节点是其父节点的左子节点，那么它的下一个节点就是它的父节点
	如果一个节点没有右子树，且它还是父节点的右子节点，那么我们可以沿着指向父节点的指针一直向上遍历，直到找到一个是它父节点的左子节点的节点，如果这样的节点存在，那么这个节点的父节点就是我们要找的下一个节点。
 */
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}

	var node *TreeLinkNode

	if pNode.Right != nil {
		rn := pNode.Right
		for rn.Left != nil {
			rn = rn.Left
		}
		node = rn
	} else {
		cur, par := pNode, pNode.Next

		for par != nil && cur == par.Right {
			cur = par
			par = par.Next
		}

		node = par
	}
	return node
}