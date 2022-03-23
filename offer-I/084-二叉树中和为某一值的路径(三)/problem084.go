/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/6 8:14 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _84_二叉树中和为某一值的路径_三_


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ans int
func FindPath( root *TreeNode ,  sum int ) int {

	if root == nil {
		return 0
	}

	dfs(root, sum, 0)
	FindPath(root.Left, sum)
	FindPath(root.Right,sum)
	return ans
}

func dfs(root *TreeNode, sum, target int) {
	if root == nil {
		return
	}
	target += root.Val
	if target == sum {
		ans++
	}
	dfs(root.Left, sum, target)
	dfs(root.Right, sum, target)
	target -= root.Val

}