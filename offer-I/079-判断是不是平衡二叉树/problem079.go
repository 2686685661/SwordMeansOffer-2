/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 2:18 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _79_判断是不是平衡二叉树

import "math"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 从上往下递归，重复遍历下层节点
func IsBalanced_Solution( pRoot *TreeNode ) bool {
	if pRoot == nil {
		return true
	}

	if math.Abs(float64(maxDepth(pRoot.Left) - maxDepth(pRoot.Right))) <= 1 {
		return IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
	} else {
		return false
	}
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := maxDepth(root.Left)
	n := maxDepth(root.Right)
	if m > n {
		return m + 1
	} else {
		return n + 1
	}
}

// 从下往上遍历
// 如果子树是平衡二叉树，则返回子树的高度；如果发现子树不是平衡二叉树，则直接停止遍历，这样至多只对每个结点访问一次。
func IsBalanced_Solution2( pRoot *TreeNode ) bool {
	if pRoot == nil {
		return true
	}
	return getDepth(pRoot) != -1
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getDepth(root.Left)
	if l == -1 {
		return -1
	}
	r := getDepth(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(float64(l - r)) > 1 {
		return -1
	} else {
		return int(math.Max(float64(l), float64(r))) + 1
	}

}