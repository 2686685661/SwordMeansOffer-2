/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/12 11:40 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _54_所有大于等于节点的值之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {

	count := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		count += node.Val
		node.Val = count
		inorder(node.Left)
	}
	inorder(root)
	return root

}