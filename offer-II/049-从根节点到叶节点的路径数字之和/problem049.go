/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/10 12:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _49_从根节点到叶节点的路径数字之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}