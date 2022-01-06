/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 11:37 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _32_从上往下打印二叉树


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PrintFromTopToBottom( root *TreeNode ) []int {

	if root == nil {
		return nil
	}
	var (
		res []int
		queue []*TreeNode
	)
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		res = append(res, node.Val)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return res
}