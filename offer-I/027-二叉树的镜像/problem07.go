/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 9:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _27_二叉树的镜像

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Mirror( pRoot *TreeNode ) *TreeNode {

	if  pRoot == nil || (pRoot.Left == nil && pRoot.Right == nil) {
		return pRoot
	}
	tmp := pRoot.Left
	pRoot.Left = pRoot.Right
	pRoot.Right = tmp

	Mirror(pRoot.Left)
	Mirror(pRoot.Right)

	return pRoot

}