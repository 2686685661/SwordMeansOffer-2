/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 3:58 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _36_二叉搜索树与双向链表

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pre *TreeNode
func Convert( pRootOfTree *TreeNode ) *TreeNode {

	if pRootOfTree == nil {
		return nil
	}
	covertCore(pRootOfTree)
	res := pRootOfTree
	for res.Left != nil {
		res = res.Left
	}
	return res

}
func covertCore(cur *TreeNode) {
	if cur == nil {
		return
	}
	covertCore(cur.Left)

	cur.Left = pre
	if pre != nil {
		pre.Right = cur
	}
	pre = cur

	covertCore(cur.Right)
}