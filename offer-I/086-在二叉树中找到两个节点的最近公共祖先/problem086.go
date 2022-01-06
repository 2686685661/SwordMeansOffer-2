/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/6 8:55 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _86_在二叉树中找到两个节点的最近公共祖先

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
	node := dfs(root, o1, o2)
	return node.Val
}

func dfs(root *TreeNode, o1, o2 int) *TreeNode {


	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}

	left := dfs(root.Left, o1, o2)
	right := dfs(root.Right, o1, o2)
	if left == nil {
		return right
	} else {
		if right == nil {
			return left
		} else {
			return root
		}
	}
}