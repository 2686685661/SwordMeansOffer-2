/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/15 2:31 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _54_二叉搜索树的第k个结点

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var count = 0
func KthNode( pRoot *TreeNode ,  k int ) *TreeNode {

	if pRoot == nil || k == 0 {
		return nil
	}
	if node := KthNode(pRoot.Left, k); node != nil {
		return node
	}
	count += 1
	if count == k {
		return pRoot
	}

	if node := KthNode(pRoot.Right, k); node != nil {
		return node
	}
	return nil
}