/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/7 1:29 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _45__二叉树最底层最左边的值

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func findBottomLeftValue(root *TreeNode) int {
	//return bfs(root)
	dfs(root, 0)
	return ans
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	ans := 0
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}


// dfs
var (
	maxDepth = -1
	ans = 0
)
func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if depth > maxDepth {
		ans = root.Val
		maxDepth = depth
	}

	dfs(root.Left, depth+1)
	dfs(root.Right, depth+1)

}
