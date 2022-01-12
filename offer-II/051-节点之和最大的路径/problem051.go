/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/11 12:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _51_节点之和最大的路径

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// https://leetcode-cn.com/problems/jC7MId/solution/jian-zhi-offer-2-mian-shi-ti-51-shu-zhon-aumb/

// dfs 后续遍历
func maxPathSum(root *TreeNode) int {
	var (
		ret = root.Val
		dfs func(node *TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := MaxInt(0, dfs(node.Left))
		right := MaxInt(0, dfs(node.Right))
		ret = MaxInt(ret, left + right + node.Val)
		return node.Val + MaxInt(left, right)
	}
	dfs(root)
	return ret
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}