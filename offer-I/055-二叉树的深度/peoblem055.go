/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 3:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _55_二叉树的深度


type TreeNode struct {
	Val int
	Left *TreeNode
     Right *TreeNode
}

func TreeDepth( pRoot *TreeNode ) int {
	if pRoot == nil {
		return 0
	}
	m := TreeDepth(pRoot.Left)
	n := TreeDepth(pRoot.Right)
	if m > n {
		return m + 1
	} else {
		return n + 1
	}
}