/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 8:25 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _26_树的子结构

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func HasSubtree( pRoot1 *TreeNode ,  pRoot2 *TreeNode ) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}

	return isSubtree(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func isSubtree(pRoot1 *TreeNode ,  pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return isSubtree(pRoot1.Left, pRoot2.Left) && isSubtree(pRoot1.Right, pRoot2.Right)
}