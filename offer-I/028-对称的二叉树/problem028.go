/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 9:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _28_对称的二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetrical( pRoot *TreeNode ) bool {
	if pRoot == nil {
		return true
	}
	return isSymmetricalCore(pRoot.Left, pRoot.Right)
}

func isSymmetricalCore(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val {
		return isSymmetricalCore(left.Left, right.Right) && isSymmetricalCore(left.Right, right.Left)
	}
	return false
}