/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/12 11:09 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _52_展平二叉搜索树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {

	tmp := &TreeNode{}
	cur := tmp
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		cur.Right = node
		node.Left = nil
		cur = node

		inorder(node.Right)
	}

	inorder(root)
	return tmp.Right

}