/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 7:48 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _82_二叉树中和为某一值的路径_一_

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum( root *TreeNode ,  sum int ) bool {
	if root == nil {
		return false
	}
	sum -= root.Val

	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
