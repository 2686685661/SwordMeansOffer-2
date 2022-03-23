/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/12 11:25 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _53_二叉搜索树中的中序后继

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 我tm反手一个中序遍历+数组
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node)
		dfs(node.Right)
	}
	dfs(root)

	for i, node := range res {
		if node == p && i+1 < len(res) {
			return res[i+1]
		}
	}
	return nil
}

// 中序遍历记录上一个节点pre，若pre==p，那么当前节点node就是p的中序后继
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {

	if root == nil || p == nil {
		return nil
	}

	var dfs func(node *TreeNode)
	var pre, res *TreeNode
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == p {
			res = node
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return res
}
