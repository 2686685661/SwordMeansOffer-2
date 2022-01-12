/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/10 11:20 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _47_二叉树剪枝

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := pruneTree(root.Left), pruneTree(root.Right)
	if root.Val == 0 && left == nil && right == nil {
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}