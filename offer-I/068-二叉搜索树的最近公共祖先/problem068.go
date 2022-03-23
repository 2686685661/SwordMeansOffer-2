/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 7:48 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _68_二叉搜索树的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor( root *TreeNode ,  p int ,  q int ) int {
	if root == nil {
		return -1
	}
	if root.Val == p || root.Val == q {
		return root.Val
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != -1 && right != -1 {
		return root.Val
	}
	if left != -1 {
		return left
	}
	if right != -1 {
		return right
	}
	return -1
}


func lowestCommonAncestor2( root *TreeNode ,  p int ,  q int ) int {
	if root.Val < p && root.Val < q {
		return lowestCommonAncestor2(root.Right, p, q)
	} else if root.Val > p && root.Val > q {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	return root.Val
}

